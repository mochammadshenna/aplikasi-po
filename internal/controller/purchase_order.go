package controller

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mochammadshenna/aplikasi-po/internal/model/api"
	"github.com/mochammadshenna/aplikasi-po/internal/service"
)

// PurchaseOrderController interface defines all the handler methods
type PurchaseOrderController interface {
	// Auth handlers
	HandleLoginPage(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error

	// Page handlers
	HandlePurchaseOrdersPage(c *fiber.Ctx) error
	HandlePurchaseOrderForm(c *fiber.Ctx) error

	// API handlers
	HandleGetPurchaseOrders(c *fiber.Ctx) error
	HandleGetPurchaseOrderById(c *fiber.Ctx) error
	HandleSavePurchaseOrder(c *fiber.Ctx) error
	HandleUpdatePurchaseOrder(c *fiber.Ctx) error
	HandleDeletePurchaseOrder(c *fiber.Ctx) error

	// Factory handlers
	HandleGetProductionFactories(c *fiber.Ctx) error
	HandleGetFinishingFactories(c *fiber.Ctx) error

	// Logout handler
	HandleLogout(c *fiber.Ctx) error

	// Dashboard handler
	HandleDashboardPage(c *fiber.Ctx) error
}

// PurchaseOrderControllerImpl implements the PurchaseOrderController interface
type PurchaseOrderControllerImpl struct {
	PurchaseOrderService service.Service
}

// NewPurchaseOrderController creates a new instance of PurchaseOrderController
func NewPurchaseOrderController(svc service.Service) PurchaseOrderController {
	return &PurchaseOrderControllerImpl{
		PurchaseOrderService: svc,
	}
}

// Implementation of all interface methods
func (controller *PurchaseOrderControllerImpl) HandleLoginPage(c *fiber.Ctx) error {
	log.Println("Rendering login page...")
	return c.Render("pages/login", fiber.Map{
		"Title": "Login - Purchase Order System",
	})
}

func (controller *PurchaseOrderControllerImpl) Login(c *fiber.Ctx) error {
	var request api.AuthAdminRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	response, err := controller.PurchaseOrderService.Login(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Set cookie for web clients
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = response.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true
	c.Cookie(cookie)

	// Return response based on Accept header
	if c.Get("Accept") == "application/json" {
		return c.JSON(response)
	}

	// Redirect to dashboard for web clients
	return c.Redirect("/dashboard")
}

func (controller *PurchaseOrderControllerImpl) GoogleLogin(c *fiber.Ctx) error {
	credential := c.FormValue("credential")
	response, err := controller.PurchaseOrderService.GoogleLogin(c.Context(), credential)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}

func (controller *PurchaseOrderControllerImpl) HandlePurchaseOrdersPage(c *fiber.Ctx) error {
	orders, err := controller.PurchaseOrderService.FindAllPurchaseOrder(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load purchase orders",
		})
	}

	return c.Render("pages/purchase-orders", fiber.Map{
		"Title": "Purchase Orders",
		"List":  orders.List,
	})
}

func (controller *PurchaseOrderControllerImpl) HandlePurchaseOrderForm(c *fiber.Ctx) error {
	var data fiber.Map = fiber.Map{}

	if id := c.Params("id"); id != "" {
		poId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID format",
			})
		}

		poRequest := api.FindPurchaseOrderByIdRequest{Id: poId}
		po := controller.PurchaseOrderService.FindPurchaseOrderById(c.Context(), poRequest)
		if po.Id == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Purchase order not found",
			})
		}

		data = fiber.Map{
			"Id":                po.Id,
			"ProductionFactory": po.ProductionFactory,
			"PICName":           po.PICName,
			"QuantityPO":        po.QuantityPO,
			"PaymentTerm":       po.PaymentTerm,
			"ProductItem":       po.ProductItem,
			"Status":            po.Status,
			"CreatedAt":         po.CreatedAt,
		}
	}

	// Get factories for dropdowns
	factories := controller.PurchaseOrderService.FindProductionFactory(c.Context(), api.FindFactoryByIdRequest{})
	data["ProductionFactories"] = factories

	return c.Render("partials/purchase-order-form", data)
}

func (controller *PurchaseOrderControllerImpl) HandleGetPurchaseOrders(c *fiber.Ctx) error {
	orders, err := controller.PurchaseOrderService.FindAllPurchaseOrder(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load purchase orders",
		})
	}

	if c.Get("HX-Request") == "true" {
		return c.Render("partials/purchase-order-rows", fiber.Map{
			"List": orders.List,
		})
	}

	return c.JSON(orders)
}

func (controller *PurchaseOrderControllerImpl) HandleGetPurchaseOrderById(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	request := api.FindPurchaseOrderByIdRequest{Id: id}
	response := controller.PurchaseOrderService.FindPurchaseOrderById(c.Context(), request)
	if response.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Purchase order not found",
		})
	}

	return c.JSON(response)
}

func (controller *PurchaseOrderControllerImpl) HandleSavePurchaseOrder(c *fiber.Ctx) error {
	var request api.SavePurchaseOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	response, err := controller.PurchaseOrderService.SavePurchaseOrder(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if c.Get("HX-Request") == "true" {
		return c.Render("partials/alert", fiber.Map{
			"Type":    "success",
			"Message": "Purchase order created successfully",
		})
	}

	return c.JSON(response)
}

func (controller *PurchaseOrderControllerImpl) HandleUpdatePurchaseOrder(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var request api.UpdatePurchaseOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}
	request.Id = id

	response, err := controller.PurchaseOrderService.UpdatePurchaseOrder(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if c.Get("HX-Request") == "true" {
		return c.Render("partials/alert", fiber.Map{
			"Type":    "success",
			"Message": "Purchase order updated successfully",
		})
	}

	return c.JSON(response)
}

func (controller *PurchaseOrderControllerImpl) HandleDeletePurchaseOrder(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	request := api.DeletePurchaseOrderRequest{Id: id}
	_, err = controller.PurchaseOrderService.DeletePurchaseOrder(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if c.Get("HX-Request") == "true" {
		return c.Render("partials/alert", fiber.Map{
			"Type":    "success",
			"Message": "Purchase order deleted successfully",
		})
	}

	return c.JSON(fiber.Map{"message": "Purchase order deleted successfully"})
}

func (controller *PurchaseOrderControllerImpl) HandleGetProductionFactories(c *fiber.Ctx) error {
	factories := controller.PurchaseOrderService.FindProductionFactory(c.Context(), api.FindFactoryByIdRequest{})

	return c.JSON(factories)
}

func (controller *PurchaseOrderControllerImpl) HandleGetFinishingFactories(c *fiber.Ctx) error {
	factories := controller.PurchaseOrderService.FindFinishingFactory(c.Context(), api.FindFactoryByIdRequest{})

	return c.JSON(factories)
}

func (controller *PurchaseOrderControllerImpl) HandleLogout(c *fiber.Ctx) error {
	// Clear the JWT cookie
	c.ClearCookie("jwt")

	// Redirect to login page
	return c.Redirect("/")
}

func (controller *PurchaseOrderControllerImpl) HandleDashboardPage(c *fiber.Ctx) error {
	return c.Render("pages/dashboard", fiber.Map{
		"Title": "Dashboard",
	})
}
