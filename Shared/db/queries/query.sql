--   ____                  _
--  / __ \__  _____  _____(_)__  _____
-- / / / / / / / _ \/ ___/ / _ \/ ___/
--/ /_/ / /_/ /  __/ /  / /  __(__  )
--\___\_\__,_/\___/_/  /_/\___/____/
--
-- This File Contains all Queries on the Main Database.
-- Refer to Sqlc for more information https://docs.sqlc.dev/en/stable/
--
-- The File Includes 5 Section:
-- 1- `Selections` following: Get? / Get?By?
-- 2- `Creation` following Create?: For stand alone Element
-- 3- `Adding` following Add?To?: For Joins
-- 4- `Updates` following: Update?
-- 5- `Deletions` following: Delete?
--
-- Rq: Please use PascalCase for naming.

------------------
-- 1- Selections
------------------

-- name: GetBlock :many
select id, author, created_at, updated_at from blocks;

-- name: GetBlockByID :one
select b.version_number
from blocks b
where b.id = $1;

-- name: GetBlockImages :many
select bi.title, bi.alt, bi.file
from block_images bi
inner join block_langs bl
on block_lang.id = bi.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockText :many
select bt.content, bt.hint, bt.name
from block_texts bt
inner join block_langs bl
on bl.id = bt.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockRichText :many
select brt.name, brt.hint, brt.content
from block_rich_texts brt
inner join block_langs bl
on bl.id = brt.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockCategories :many
SELECT name, descr
  FROM categories c
  INNER JOIN block_categ bc
  ON c.id = bc.categ_id
  INNER JOIN blocks
  ON bc.block_id = block.id;

-- name: GetBlockTags :many
SELECT name, descr
  FROM categories c
  INNER JOIN block_tags bt
  ON c.id = bt.tag_id
  INNER JOIN blocks
  ON bt.block_id = block.id;

------------------
-- 2- Creations
------------------

-- name: CreateBlock :exec
INSERT INTO blocks (
    has_likes, has_comments, block_type, comments_type
) VALUES ($1,$2,$3,$4);

-- name: CreateTag :exec
INSERT INTO Tags (
    name, descr
) VALUES ($1,$2);

-- name: CreateCateg :exec
INSERT INTO Tags (
    name, descr
) VALUES ($1,$2);

------------------
-- 3- Adding
------------------

-- name: AddTagToBlock :exec
INSERT INTO block_tags(
    block_id, tag_id
) VALUES ($1,$2);

-- name: AddCategToBlock :exec
INSERT INTO block_categ(
    block_id, categ_id
) VALUES ($1,$2);

-- name: AddCommentRules :exec
INSERT INTO comment_types(
    nested, has_likes, editable, max_nest
) VALUES ($1,$2,$3,$4) RETURNING id;

------------------
-- 5- Updates
------------------

------------------
-- 5- Deletions
------------------

-- name: DeleteBlock :exec
DELETE FROM blocks
    WHERE id = $1
    RETURNING id;

-- name: DeleteBlockLang :exec
DELETE FROM block_langs
    WHERE block_id = $1
    AND lang_name = $2
    RETURNING id;
