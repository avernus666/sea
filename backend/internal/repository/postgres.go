package repository

import (
 "context"
 "encoding/json"
 "errors"
 "sea/backend/internal/domain"
 "github.com/jackc/pgx/v5"
 "github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct { pool *pgxpool.Pool }
func NewPostgresRepository(ctx context.Context,url string)(*PostgresRepository,error){
 pool,err:=pgxpool.New(ctx,url); if err!=nil{return nil,err}
 if err=pool.Ping(ctx);err!=nil{pool.Close();return nil,err}
 return &PostgresRepository{pool:pool},nil
}
func(r *PostgresRepository)Close(){r.pool.Close()}
func(r *PostgresRepository)List(ctx context.Context,q string)([]domain.Institution,error){
 rows,err:=r.pool.Query(ctx,`SELECT id,name,city,kind,website,description,tags FROM institutions WHERE $1='' OR name ILIKE '%'||$1||'%' OR city ILIKE '%'||$1||'%' OR kind ILIKE '%'||$1||'%' ORDER BY featured DESC,name ASC`,q)
 if err!=nil{return nil,err};defer rows.Close()
 items:=make([]domain.Institution,0)
 for rows.Next(){var x domain.Institution;var tags []byte;if err:=rows.Scan(&x.ID,&x.Name,&x.City,&x.Kind,&x.Website,&x.Description,&tags);err!=nil{return nil,err};if err:=json.Unmarshal(tags,&x.Tags);err!=nil{return nil,err};items=append(items,x)}
 return items,rows.Err()
}
func(r *PostgresRepository)GetByID(ctx context.Context,id string)(domain.Institution,error){
 var x domain.Institution;var tags []byte
 err:=r.pool.QueryRow(ctx,`SELECT id,name,city,kind,website,description,tags FROM institutions WHERE id=$1`,id).Scan(&x.ID,&x.Name,&x.City,&x.Kind,&x.Website,&x.Description,&tags)
 if errors.Is(err,pgx.ErrNoRows){return domain.Institution{},ErrNotFound};if err!=nil{return domain.Institution{},err}
 if err=json.Unmarshal(tags,&x.Tags);err!=nil{return domain.Institution{},err};return x,nil
}
