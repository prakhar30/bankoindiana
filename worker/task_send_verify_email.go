package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	db "github.com/prakhar30/bankoindiana/db/sqlc"
	"github.com/prakhar30/bankoindiana/utils"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("could not marshal payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("could not enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task: send_verify_email")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exist %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get user %w", err)
	}

	verifyEmail, err := processor.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: utils.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify email %w", err)
	}

	// TODO: send email to user.Email with verification link
	// Imagine there exisits a function sendEmail with the processor to send email.
	// For this project, we will skip the email sending part and pick up the secret code directly from the DB.
	// Or attaching it in the log here. This is just for the sake of simplicity.

	log.Info().Str("type", task.Type()).Int("secret_code_id", int(verifyEmail.ID)).Str("secret_code", verifyEmail.SecretCode).Bytes("payload", task.Payload()).Str("username", user.Username).Str("email", user.Email).Msg("sending verification email")
	return nil
}
