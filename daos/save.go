package storage

import (
    "context"
    "Farmako/config"
    "Farmako/model"

    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type CouponStore struct {
    db *gorm.DB
}


func NewCouponStore() *CouponStore {
    return &CouponStore{db: config.DB}
}


func (s *CouponStore) Add(ctx context.Context, c *model.Coupon) error {
    return s.db.WithContext(ctx).Create(c).Error
}

func (s *CouponStore) FindByCode(ctx context.Context, code string) (*model.Coupon, error) {
    var c model.Coupon
    err := s.db.WithContext(ctx).
        Where("code = ?", code).
        First(&c).Error
    if err != nil {
        return nil, err
    }
    return &c, nil
}


func (s *CouponStore) Use(ctx context.Context, code string) error {
    return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        var c model.Coupon
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            Where("code = ?", code).
            First(&c).Error; err != nil {
            return err
        }
      
        return nil
    })
}
