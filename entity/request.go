package entity

type CheckPersonExistByTelegramIDRequest struct {
	TelegramID int64 `json:"telegramId,omitempty"  validate:"required"`
}
