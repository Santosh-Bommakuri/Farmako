package handler

import (
    "Farmako/model"
    "Farmako/service"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type CouponHandler struct {
    Service *service.CouponService
}


func (h *CouponHandler) CreateCoupon(c *gin.Context) {
    var in model.Coupon
    if err := c.BindJSON(&in); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    ctx := c.Request.Context()
    if err := h.Service.AddCoupon(ctx, &in); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, in)
}


func (h *CouponHandler) ApplyCoupon(c *gin.Context) {
    var req struct {
        Code       string     `json:"code"`
        OrderTotal float64    `json:"order_total"`
        Time       *time.Time `json:"time"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    now := time.Now()
    if req.Time != nil {
        now = *req.Time
    }
    ctx := c.Request.Context()
    discount, err := h.Service.ApplyCoupon(ctx, req.Code, req.OrderTotal, now)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": req.Code, "discount": discount})
}
