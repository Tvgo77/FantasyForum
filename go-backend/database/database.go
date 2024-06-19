package database

import (
	"context"
	"go-backend/domain"
	"log"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

func NewDatabase(dsn string) (domain.Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &postgresDB{db: db}, nil
}

func NewDatabaseFromExist(db *gorm.DB) domain.Database{
	return &postgresDB{db: db}
}

func (p *postgresDB) AutoMigrate(dest ...interface{}) error {
	return p.db.AutoMigrate(dest...)
}

func (p *postgresDB) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return p.db.Begin(opts...)
}

func (p *postgresDB) SavePoint(name string) *gorm.DB {
	return p.db.SavePoint(name)
}

func (p *postgresDB) Rollbackto(name string) *gorm.DB {
	return p.db.RollbackTo(name)
}

func (p *postgresDB) Rollback() *gorm.DB {
	return p.db.Rollback()
}

func (p *postgresDB) Create(value interface{}) (tx *gorm.DB) {
	return p.db.Create(value)
}

func (p *postgresDB) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return p.db.First(dest, conds...)
}

func (p *postgresDB) Select(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return p.db.Select(query, args...)
}

func (p *postgresDB) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return p.db.Where(query, args...)
}

func (p *postgresDB) WithContext(ctx context.Context) (tx *gorm.DB) {
	return p.db.WithContext(ctx)
}

/* Custom CRUD convinent interface */
// Note: Only non-zero of struct field will be used for condition and updates
// If you want to deal with zero value, use map[string]interface{}{} as input args

func (p *postgresDB) Ping() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Should insert struct to table and assign back to default primary key ID
func (p *postgresDB) InsertOne(ctx context.Context, src interface{}) error {
	result := p.db.WithContext(ctx).Create(src)
	return result.Error
}

// Query row goes into dest
func (p *postgresDB) FindOne(ctx context.Context, dest interface{}, conds interface{}) error {
	result := p.db.WithContext(ctx).Where(conds).First(dest)
	return result.Error
}


// Old one should contain primary key field
func (p *postgresDB) UpdateOne(ctx context.Context, old interface{}, new interface{}) error {
	result := p.db.WithContext(ctx).Model(old).Updates(new)
	return result.Error
}

func (p *postgresDB) IncreaseOne(ctx context.Context, old interface{}, column string, n int) error {
	result := p.db.WithContext(ctx).Model(old).Update(column, gorm.Expr(column + " + ?", n))
	return result.Error
}

// Arg should contain primary key field
func (p *postgresDB) DeleteOne(ctx context.Context, arg interface{}) error {
	result := p.db.WithContext(ctx).Delete(arg)
	return result.Error
}

func (p *postgresDB) CountRows(ctx context.Context, conds interface{}) (int, error) {
	var count int64
	result := p.db.WithContext(ctx).Model(conds).Where(conds).Count(&count)
	return int(count), result.Error
}

