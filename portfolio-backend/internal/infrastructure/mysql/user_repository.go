package mysql

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/kojima1128/portfolio-backend/internal/model"
	"github.com/kojima1128/portfolio-backend/internal/repository"
)

type userRecord struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CognitoID string    `gorm:"column:cognito_id;not null"`
	Name      string    `gorm:"column:name;not null"`
	TenantID  string    `gorm:"column:tenant_id;not null"`
	SiteID    string    `gorm:"column:site_id;not null"`
	Role      string    `gorm:"column:role;not null;default:user"`
	Email     string    `gorm:"column:email;unique;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (userRecord) TableName() string { return "users" }

func toModel(r *userRecord) *model.User {
	return &model.User{
		ID:        strconv.FormatUint(uint64(r.ID), 10),
		CognitoID: r.CognitoID,
		Name:      r.Name,
		TenantID:  r.TenantID,
		SiteID:    r.SiteID,
		Role:      r.Role,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Format(time.RFC3339),
		UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
	}
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	record := &userRecord{
		CognitoID: user.CognitoID,
		Name:      user.Name,
		TenantID:  user.TenantID,
		SiteID:    user.SiteID,
		Role:      user.Role,
		Email:     user.Email,
	}
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	user.ID = strconv.FormatUint(uint64(record.ID), 10)
	user.CreatedAt = record.CreatedAt.Format(time.RFC3339)
	user.UpdatedAt = record.UpdatedAt.Format(time.RFC3339)
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uid > math.MaxUint32 {
		return nil, fmt.Errorf("invalid id: %s", id)
	}
	var record userRecord
	if err := r.db.WithContext(ctx).First(&record, uid).Error; err != nil {
		return nil, err
	}
	return toModel(&record), nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	var records []userRecord
	if err := r.db.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}
	users := make([]*model.User, len(records))
	for i, rec := range records {
		users[i] = toModel(&rec)
	}
	return users, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var record userRecord
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&record).Error; err != nil {
		return nil, err
	}
	return toModel(&record), nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	uid, err := strconv.ParseUint(user.ID, 10, 64)
	if err != nil || uid > math.MaxUint32 {
		return fmt.Errorf("invalid id: %s", user.ID)
	}
	record := &userRecord{
		ID:        uint(uid),
		CognitoID: user.CognitoID,
		Name:      user.Name,
		TenantID:  user.TenantID,
		SiteID:    user.SiteID,
		Role:      user.Role,
		Email:     user.Email,
	}
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uid > math.MaxUint32 {
		return fmt.Errorf("invalid id: %s", id)
	}
	return r.db.WithContext(ctx).Delete(&userRecord{}, uid).Error
}
