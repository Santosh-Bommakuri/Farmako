package service

import (
    "context"
    "Farmako/cache"
    "Farmako/model"
    "Farmako/daos"
    "encoding/json"
    "errors"
    "fmt"
    "time"
)


type CouponService struct {
    Store *storage.CouponStore
}


func (s *CouponService) AddCoupon(ctx context.Context, c *model.Coupon) error {
    if err := s.Store.Add(ctx, c); err != nil {
        return err
    }
    raw, _ := json.Marshal(c)
    cache.SaveCoupon(fmt.Sprintf("coupon:", c.Code), raw)
    return nil
}


func (s *CouponService) GetCoupon(ctx context.Context, code string) (*model.Coupon, error) {
    key := fmt.Sprintf("coupon:", code)
    if raw, err := cache.GetCoupon(key); err == nil {
        var c model.Coupon
        if json.Unmarshal([]byte(raw), &c) == nil {
            return &c, nil
        }
    }
    c, err := s.Store.FindByCode(ctx, code)
    if err != nil {
        return nil, err
    }

    go func() {
        raw, _ := json.Marshal(c)
        cache.SaveCoupon(key, raw)
    }()
    return c, nil
}


func (s *CouponService) ApplyCoupon(ctx context.Context, code string, total float64, now time.Time) (float64, error) {
    c, err := s.GetCoupon(ctx, code)
    if err != nil {
        return 0, errors.New("coupon not found")
    }
    if now.After(c.ExpiresAt) {
        return 0, errors.New("coupon expired")
    }
    if total < c.MinOrderValue {
        return 0, errors.New("order below minimum")
    }
    if err := s.Store.Use(ctx, code); err != nil {
        return 0, errors.New("could not reserve coupon usage")
    }
    if c.DiscountType == "percent" {
        return total * c.DiscountValue / 100, nil
    }
    return c.DiscountValue, nil
}
