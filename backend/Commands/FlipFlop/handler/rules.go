package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"

	_ "github.com/lib/pq"
)

func (c *CommandService) CreateRuleGroup(ctx context.Context, rq *pb.BlockRulesRq) (*pb.ActionResponse, error) {
	q := db.New(conn)

	if rq.GetRules() == nil {
		return nil, errors.New("NO RULES PROVIDED")
	}

	if rq.GetRules().Description == "" {
		return nil, errors.New("NO DESCRIPTION PROVIDED")
	}

	params := db.CreateRuleGroupParams{
		Name: rq.GetRules().RuleName,
		Nested: sql.NullBool{
			Bool:  rq.GetRules().Nested,
			Valid: true,
		},
		HasLikes: sql.NullBool{
			Bool:  rq.GetRules().HasLikes,
			Valid: true,
		},
		HasComments: sql.NullBool{
			Bool:  rq.GetRules().HasComments,
			Valid: true,
		},
		CommentsMaxNest: sql.NullInt16{
			Int16: int16(rq.GetRules().CommentsMaxNested),
			Valid: true,
		},
		CommentsHasLikes: sql.NullBool{
			Bool:  rq.GetRules().CommentsHasLikes,
			Valid: true,
		},
		CommentEditable: sql.NullBool{
			Bool:  rq.GetRules().CommentsEditable,
			Valid: true,
		},
		Descr: rq.GetRules().Description,
	}

	id, err := q.CreateRuleGroup(ctx, params)
	if err != nil {
		log.Print(err)
		return nil, errors.New("FAILED TO CREATE RULE")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Great Success",
	}, nil
}

func (*CommandService) UpdateRuleGroup(ctx context.Context, rq *pb.BlockRulesRq) (*pb.ActionResponse, error) {
	q := db.New(conn)

	if rq.GetRules() == nil {
		return nil, errors.New("NO RULES PROVIDED")
	}

	if rq.GetRules().Description == "" {
		return nil, errors.New("NO DESCRIPTION PROVIDED")
	}

	params := db.UpdateRuleGroupParams{
		Name: rq.GetRules().RuleName,
		Nested: sql.NullBool{
			Bool:  rq.GetRules().Nested,
			Valid: true,
		},
		HasLikes: sql.NullBool{
			Bool:  rq.GetRules().HasLikes,
			Valid: true,
		},
		HasComments: sql.NullBool{
			Bool:  rq.GetRules().HasComments,
			Valid: true,
		},
		CommentsMaxNest: sql.NullInt16{
			Int16: int16(rq.GetRules().CommentsMaxNested),
			Valid: true,
		},
		CommentsHasLikes: sql.NullBool{
			Bool:  rq.GetRules().CommentsHasLikes,
			Valid: true,
		},
		CommentEditable: sql.NullBool{
			Bool:  rq.GetRules().CommentsEditable,
			Valid: true,
		},
		Descr: rq.GetRules().Description,
	}

	if err := q.UpdateRuleGroup(ctx, params); err != nil {
		log.Print(err)
		return nil, errors.New("FAILED TO CREATE RULE")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Message:   "Great Success",
	}, nil
}

func (*CommandService) DeleteRuleGroup(ctx context.Context, rq *pb.BlockRulesRq) (*pb.ActionResponse, error) {
	q := db.New(conn)

	var ruleName string
	if rq.GetRules() == nil {
		ruleName = rq.GetRuleName()
	} else {
		ruleName = rq.GetRules().RuleName
	}

	if err := q.DeleteRuleGroup(ctx, ruleName); err != nil {
		log.Print(err)
		return nil, errors.New("FAILED TO DELETE RULE")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Message:   "Great Success",
	}, nil
}
