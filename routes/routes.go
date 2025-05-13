package routes

import (
    "Farmako/handler"
    "Farmako/service"
    "Farmako/daos"

    "github.com/gin-gonic/gin"
)


func Setup() *gin.Engine {
    r := gin.Default()
    store := storage.NewCouponStore()
    svc := &service.CouponService{Store: store}
    h := &handler.CouponHandler{Service: svc}

    r.POST("/coupons/create", h.CreateCoupon)
    r.POST("/coupons/apply", h.ApplyCoupon)

    return r
}
