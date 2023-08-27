package chatcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectChat(db *sql.DB, userId string, receipientId string) (*dbmodels.Chat, error) {
	query := `SELECT * FROM Chats
	WHERE (sender_id = ? AND receiver_id = ?)
	OR (sender_id = ? AND receiver_id = ?)`

	chatData, err := crud.SelectFromDatabase(db, "Chats", query, []interface{}{userId, receipientId, receipientId, userId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("error selecting chat: %w", err)
	} else if err != nil {
		return nil, err
	}

	chat, ok := chatData[0].(*dbmodels.Chat)
	if !ok {
		return nil, fmt.Errorf("can't assert chat type")
	}

	return chat, nil
}
