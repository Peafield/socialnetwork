package chatcontrollers

import (
	"database/sql"
	"errors"
	"fmt"
	crud "socialnetwork/pkg/db/CRUD"
	"socialnetwork/pkg/db/dbstatements"
	errorhandling "socialnetwork/pkg/errorHandling"
	"socialnetwork/pkg/models/dbmodels"
)

func SelectChat(db *sql.DB, userId string, receipientId string) (*dbmodels.Chat, error) {
	chatData, err := crud.SelectFromDatabase(db, "Chats", dbstatements.SelectChatBySenderAndRecieverIdStmt, []interface{}{userId, receipientId, receipientId, userId})
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
