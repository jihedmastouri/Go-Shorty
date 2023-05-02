package db

import (
	"context"
	"log"

	pb "github.com/shorty-io/go-shorty/queries/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlockMeta(bq *pb.BlockRequest) (*pb.BlockMeta, error) {

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)

	meta := &pb.BlockMeta{}
	if err := cusror.Decode(meta); err != nil {
		return nil, err
	}

	return meta, nil
}

func GetBlock(bq *pb.BlockRequest) (*pb.Block, error) {

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)

	block := &pb.Block{}

	if err := cusror.Decode(block); err != nil {
		log.Panic(err)
		return nil, err
	}

	return block, nil
}

func GetLanguages(bq *pb.LanguageRequest) (*pb.LanguageList, error) {

	pipeline := bson.A{
		bson.M{"$match": bson.M{"block_id": bq.GetId()}},
		bson.M{"$group": bson.M{"_id": "$lang"}},
		bson.M{"$project": bson.M{"_id": 0, "lang": "$_id"}},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	results := &pb.LanguageList{}
	if err = cursor.All(context.Background(), &results.Langs); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}

func GetVersions(bq *pb.VersionsRequest) (*pb.VersionResponse, error) {
	pipeline := bson.A{
		bson.M{"$match": bson.M{"block_id": bq.GetId(), "lang": bq.GetLang()}},
		bson.M{"$group": bson.M{"_id": "$version", "changeLog": bson.M{"$first": "$changeLog"}}},
        bson.M{"$project": bson.M{"_id": 0, "version": "$_id", "changeLog": 1}},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var verList []*pb.Ver
	if err := cursor.All(context.Background(), &verList); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &pb.VersionResponse{
		Versions: verList,
	}, nil
}

func buildQuery(bq *pb.BlockRequest) (bson.M, *options.FindOneOptions) {

	if bq.GetVersion() != 0 {
		return bson.M{
			"block_id": bq.GetId(),
			"lang":     bq.GetLang(),
			"version":  bq.GetVersion(),
		}, nil
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "version", Value: -1}})
	return bson.M{
		"block_id": bq.GetId(),
		"lang":     bq.GetLang(),
	}, opts
}
