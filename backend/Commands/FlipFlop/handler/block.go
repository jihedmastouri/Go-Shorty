package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"

	_ "github.com/lib/pq"
)

type CommandService struct {
	pb.UnimplementedFlipFlopServer
}

func (c *CommandService) CreateBlock(ctx context.Context, rq *pb.CreateRequest) (*pb.ActionResponse, error) {
	q := db.New(conn)

	log.Print("Creating block:", rq.Author)

	author, err := uuid.Parse(rq.Author)
	if err != nil {
		log.Print("Failed to parse author UUID:", err)
		return nil, errors.New("FAILED TO PARSE AUTHOR ID")
	}

	rules := getBlockRules(q, rq.GetRules())

	blockType, err := q.GetTypeByName(ctx, rq.BlockType)
	if err != nil {
		log.Print("Failed to get block type:", err)
		return nil, errors.New("FAILED TO GET BLOCK TYPE")
	}

	params := db.CreateBlockParams{
		Author:           author,
		Name:             rq.Name,
		Nested:           rules.GetNested(),
		HasLikes:         rules.GetHasLikes(),
		HasComments:      rules.GetHasComments(),
		CommentsMaxNest:  int16(rules.GetCommentsMaxNested()),
		CommentsHasLikes: rules.GetCommentsHasLikes(),
		CommentEditable:  rules.GetCommentsEditable(),
		RulesName:        sql.NullString{String: rules.RuleName, Valid: true},
		Description:      sql.NullString{String: rq.Description, Valid: true},
		Type:             blockType,
	}

	id, err := q.CreateBlock(ctx, params)
	if err != nil {
		log.Print("Failed to create block:", err)
		return nil, errors.New("FAILED TO CREATE BLOCK")
	}

	langParam := db.CreateLangParams{
		LangName: "English",
		LangCode: "en_US",
		BlockID:  id,
	}

	_, err = q.CreateLang(ctx, langParam)
	if err != nil {
		log.Print("Failed to create block:", err)
		return nil, errors.New("FAILED TO CREATE BLOCK")
	}

	defer publishEvent(Msg{
		Id:        id.String(),
		LangCode:  "en_US",
		ChangeLog: "Created Block",
	}, "BlockUpdated")

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Created With Default Language en_US",
	}, nil
}

func (c *CommandService) DeleteBlock(ctx context.Context, rq *pb.DeleteRequest) (*pb.ActionResponse, error) {
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	if err = q.DeleteBlock(ctx, id); err != nil {
		log.Print("Failed to delete block:", err)
		return nil, errors.New("FAILED TO DELETE BLOCK")
	}

	publishEvent(Msg{
		Id:        rq.GetId(),
		LangCode:  "",
		ChangeLog: "Deleted Block",
	}, "BlockDeleted")

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Deleted successfully",
	}, nil
}

func (c *CommandService) CreateBlockLang(ctx context.Context, rq *pb.CreateLangRequest) (*pb.ActionResponse, error) {
	q := db.New(conn)

	blockid, err := uuid.Parse(rq.BlockId)
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	params := db.CreateLangParams{
		LangName: rq.LangName,
		LangCode: rq.LangCode,
		BlockID:  blockid,
	}

	id, err := q.CreateLang(ctx, params)
	if err != nil {
		log.Print("Failed to create block:", err)
		return nil, errors.New("FAILED TO CREATE BLOCK")
	}

	publishEvent(Msg{
		Id:        rq.BlockId,
		LangCode:  rq.LangCode,
		ChangeLog: "Created BLOCKLANG",
	}, "BlockUpdated")

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        strconv.Itoa(int(id)),
		Message:   "Created successfully",
	}, nil
}

func (c *CommandService) DeleteBlockLang(ctx context.Context, rq *pb.DeleteLangRequest) (*pb.ActionResponse, error) {
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	params := db.DeleteBlockLangParams{
		BlockID:  id,
		LangCode: rq.LangCode,
	}

	if err = q.DeleteBlockLang(ctx, params); err != nil {
		log.Print("Failed to delete block lang:", err)
		return nil, errors.New("FAILED TO DELETE BLOCK LANG")
	}

	publishEvent(Msg{
		Id:        id.String(),
		LangCode:  rq.LangCode,
		ChangeLog: "DELETED BLOCKLANG",
	}, "BlockUpdated")

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Deleted successfully",
	}, nil
}
