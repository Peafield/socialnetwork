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

func SelectGroupChat(db *sql.DB, userId string, groupId string) (*dbmodels.Chat, error) {
	chatData, err := crud.SelectFromDatabase(db, "Chats", dbstatements.SelectGroupChatStmt, []interface{}{groupId})
	if err != nil && !errors.Is(err, errorhandling.ErrNoResultsFound) {
		return nil, fmt.Errorf("error selecting group chat: %w", err)
	} else if err != nil {
		return nil, err
	}

	chat, ok := chatData[0].(*dbmodels.Chat)
	if !ok {
		return nil, fmt.Errorf("can't assert chat type")
	}

	return chat, nil
}
