// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: taxonomy.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addCategToBlock = `-- name: AddCategToBlock :exec
INSERT INTO block_categs(
    block_id, categ_id
) VALUES ($1, (SELECT id FROM categories WHERE name = $2))
`

type AddCategToBlockParams struct {
	BlockID uuid.UUID
	Name    string
}

func (q *Queries) AddCategToBlock(ctx context.Context, arg AddCategToBlockParams) error {
	_, err := q.db.ExecContext(ctx, addCategToBlock, arg.BlockID, arg.Name)
	return err
}

const addTagToBlock = `-- name: AddTagToBlock :exec
INSERT INTO block_tags(
    block_id, tag_id
) VALUES ($1, (SELECT id FROM tags WHERE name = $2))
`

type AddTagToBlockParams struct {
	BlockID uuid.UUID
	Name    string
}

func (q *Queries) AddTagToBlock(ctx context.Context, arg AddTagToBlockParams) error {
	_, err := q.db.ExecContext(ctx, addTagToBlock, arg.BlockID, arg.Name)
	return err
}

const createCateg = `-- name: CreateCateg :one
INSERT INTO categories (
    name, descr
) VALUES ($1,$2) RETURNING name
`

type CreateCategParams struct {
	Name  string
	Descr sql.NullString
}

func (q *Queries) CreateCateg(ctx context.Context, arg CreateCategParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createCateg, arg.Name, arg.Descr)
	var name string
	err := row.Scan(&name)
	return name, err
}

const createTag = `-- name: CreateTag :one
INSERT INTO tags (
    name, descr
) VALUES ($1,$2) RETURNING name
`

type CreateTagParams struct {
	Name  string
	Descr sql.NullString
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createTag, arg.Name, arg.Descr)
	var name string
	err := row.Scan(&name)
	return name, err
}

const deleteAllBlockCategs = `-- name: DeleteAllBlockCategs :exec
DELETE FROM block_categs
WHERE block_id = $1
`

func (q *Queries) DeleteAllBlockCategs(ctx context.Context, blockID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllBlockCategs, blockID)
	return err
}

const deleteAllBlockTags = `-- name: DeleteAllBlockTags :exec
DELETE FROM block_tags
WHERE block_id = $1
`

func (q *Queries) DeleteAllBlockTags(ctx context.Context, blockID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllBlockTags, blockID)
	return err
}

const deleteBlockCateg = `-- name: DeleteBlockCateg :exec
DELETE FROM block_categs
WHERE block_id = $1 AND
    categ_id = (
        SELECT id
        FROM categories
        WHERE name = $2
    )
`

type DeleteBlockCategParams struct {
	BlockID uuid.UUID
	Name    string
}

func (q *Queries) DeleteBlockCateg(ctx context.Context, arg DeleteBlockCategParams) error {
	_, err := q.db.ExecContext(ctx, deleteBlockCateg, arg.BlockID, arg.Name)
	return err
}

const deleteBlockTag = `-- name: DeleteBlockTag :exec
DELETE FROM block_tags
WHERE block_id = $1 AND
    tag_id = (
        SELECT id
        FROM tags
        WHERE name = $2
    )
`

type DeleteBlockTagParams struct {
	BlockID uuid.UUID
	Name    string
}

func (q *Queries) DeleteBlockTag(ctx context.Context, arg DeleteBlockTagParams) error {
	_, err := q.db.ExecContext(ctx, deleteBlockTag, arg.BlockID, arg.Name)
	return err
}

const deleteCateg = `-- name: DeleteCateg :one
DELETE FROM categories
WHERE name = $1
RETURNING id
`

func (q *Queries) DeleteCateg(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRowContext(ctx, deleteCateg, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteCategByID = `-- name: DeleteCategByID :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategByID, id)
	return err
}

const deleteTag = `-- name: DeleteTag :one
DELETE FROM tags
WHERE name = $1
RETURNING id
`

func (q *Queries) DeleteTag(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRowContext(ctx, deleteTag, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteTagByID = `-- name: DeleteTagByID :exec
DELETE FROM tags
WHERE id = $1
`

func (q *Queries) DeleteTagByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTagByID, id)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT name, descr
FROM categories
LIMIT 100
`

type GetAllCategoriesRow struct {
	Name  string
	Descr sql.NullString
}

func (q *Queries) GetAllCategories(ctx context.Context) ([]GetAllCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllCategoriesRow
	for rows.Next() {
		var i GetAllCategoriesRow
		if err := rows.Scan(&i.Name, &i.Descr); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTags = `-- name: GetAllTags :many
SELECT name, descr
FROM tags
LIMIT 100
`

type GetAllTagsRow struct {
	Name  string
	Descr sql.NullString
}

func (q *Queries) GetAllTags(ctx context.Context) ([]GetAllTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTagsRow
	for rows.Next() {
		var i GetAllTagsRow
		if err := rows.Scan(&i.Name, &i.Descr); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBlockCategories = `-- name: GetBlockCategories :many
SELECT c.name
FROM categories c
INNER JOIN block_categs bc
ON c.id = bc.categ_id
WHERE bc.block_id = $1
LIMIT 100
`

func (q *Queries) GetBlockCategories(ctx context.Context, blockID uuid.UUID) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getBlockCategories, blockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBlockTags = `-- name: GetBlockTags :many
SELECT t.name
FROM tags t
INNER JOIN block_tags bt
ON t.id = bt.tag_id
WHERE bt.block_id = $1
LIMIT 100
`

func (q *Queries) GetBlockTags(ctx context.Context, blockID uuid.UUID) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getBlockTags, blockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
Update categories
    SET name = $1,
        descr = $2
WHERE name = $1
RETURNING id
`

type UpdateCategoryParams struct {
	Name  string
	Descr sql.NullString
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.Name, arg.Descr)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateCategoryById = `-- name: UpdateCategoryById :exec
Update categories
    SET name = $2,
        descr = $3
WHERE id = $1
`

type UpdateCategoryByIdParams struct {
	ID    int32
	Name  string
	Descr sql.NullString
}

func (q *Queries) UpdateCategoryById(ctx context.Context, arg UpdateCategoryByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateCategoryById, arg.ID, arg.Name, arg.Descr)
	return err
}

const updateTag = `-- name: UpdateTag :one
Update tags
    SET name = $1,
        descr = $2
WHERE name = $1
RETURNING id
`

type UpdateTagParams struct {
	Name  string
	Descr sql.NullString
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateTag, arg.Name, arg.Descr)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateTagById = `-- name: UpdateTagById :exec
Update tags
    SET name = $2,
        descr = $3
WHERE id = $1
`

type UpdateTagByIdParams struct {
	ID    int32
	Name  string
	Descr sql.NullString
}

func (q *Queries) UpdateTagById(ctx context.Context, arg UpdateTagByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateTagById, arg.ID, arg.Name, arg.Descr)
	return err
}
