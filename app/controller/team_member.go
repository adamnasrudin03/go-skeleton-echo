package controller

import (
	"net/http"
	"strconv"
	"strings"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-skeleton-echo/app/configs"
	"github.com/adamnasrudin03/go-skeleton-echo/app/dto"
	"github.com/adamnasrudin03/go-skeleton-echo/app/middlewares"
	"github.com/adamnasrudin03/go-skeleton-echo/app/service"
	"github.com/adamnasrudin03/go-skeleton-echo/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type TeamMemberController interface {
	Mount(group *echo.Group)
	Create(c echo.Context) error
	GetDetail(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
	GetList(c echo.Context) error
}

type TeamMemberHandler struct {
	Service  service.TeamMemberService
	Cfg      *configs.Configs
	Logger   *logrus.Logger
	Validate *validator.Validate
}

func NewTeamMemberDelivery(
	srv service.TeamMemberService,
	cfg *configs.Configs,
	logger *logrus.Logger,
	validator *validator.Validate,
) TeamMemberController {
	return &TeamMemberHandler{
		Service:  srv,
		Cfg:      cfg,
		Logger:   logger,
		Validate: validator,
	}
}
func (h *TeamMemberHandler) Mount(group *echo.Group) {
	group.POST("", h.Create, middlewares.BasicAuth(h.Cfg.App.BasicUsername, h.Cfg.App.BasicPassword))
	group.GET("", h.GetList)
	group.GET("/:id", h.GetDetail)
	group.DELETE("/:id", h.Delete, middlewares.BasicAuth(h.Cfg.App.BasicUsername, h.Cfg.App.BasicPassword))
	group.PUT("/:id", h.Update, middlewares.BasicAuth(h.Cfg.App.BasicUsername, h.Cfg.App.BasicPassword))
}

func (h *TeamMemberHandler) Create(c echo.Context) error {
	var (
		opName = "TeamMemberController-Create"
		ctx    = c.Request().Context()
		input  dto.TeamMemberCreateReq
		err    error
	)

	err = c.Bind(&input)
	if err != nil {
		h.Logger.Errorf("%v error bind json: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrGetRequest())
	}

	// validation input user
	err = h.Validate.Struct(input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, response_mapper.FormatValidationError(err))
	}

	res, err := h.Service.Create(ctx, input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.JSON(http.StatusCreated, response_mapper.RenderStruct(http.StatusCreated, res))
}

func (h *TeamMemberHandler) GetDetail(c echo.Context) error {
	var (
		opName  = "TeamMemberController-GetDetail"
		ctx     = c.Request().Context()
		idParam = strings.TrimSpace(c.Param("id"))
		err     error
	)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.Logger.Errorf("%v error parse param: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrInvalid("ID Anggota team", "Team Member ID"))
	}

	res, err := h.Service.GetByID(ctx, id)
	if err != nil {
		return utils.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, response_mapper.RenderStruct(http.StatusOK, res))
}

func (h *TeamMemberHandler) Delete(c echo.Context) error {
	var (
		opName  = "TeamMemberController-Delete"
		ctx     = c.Request().Context()
		idParam = strings.TrimSpace(c.Param("id"))
		err     error
	)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.Logger.Errorf("%v error parse param: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrInvalid("ID Anggota team", "Team Member ID"))
	}

	err = h.Service.DeleteByID(ctx, id)
	if err != nil {
		return utils.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, response_mapper.RenderStruct(http.StatusOK, response_mapper.MultiLanguages{
		ID: "Anggota Tim Berhasil Dihapus",
		EN: "Team Member Deleted Successfully",
	}))
}

func (h *TeamMemberHandler) Update(c echo.Context) error {
	var (
		opName  = "TeamMemberController-Update"
		ctx     = c.Request().Context()
		idParam = strings.TrimSpace(c.Param("id"))
		input   dto.TeamMemberUpdateReq
		err     error
	)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.Logger.Errorf("%v error parse param: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrInvalid("ID Anggota team", "Team Member ID"))
	}

	err = c.Bind(&input)
	if err != nil {
		h.Logger.Errorf("%v error bind json: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrGetRequest())
	}
	input.ID = id

	// validation input user
	err = h.Validate.Struct(input)
	if err != nil {
		return utils.HttpError(c, response_mapper.FormatValidationError(err))
	}

	err = h.Service.Update(ctx, input)
	if err != nil {
		return utils.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, response_mapper.RenderStruct(http.StatusOK, response_mapper.MultiLanguages{
		ID: "Anggota Tim Berhasil Diperbarui",
		EN: "Team Member Updated Successfully",
	}))
}

func (h *TeamMemberHandler) GetList(c echo.Context) error {
	var (
		opName = "TeamMemberController-GetList"
		ctx    = c.Request().Context()
		input  dto.TeamMemberListReq
		err    error
	)

	err = c.Bind(&input)
	if err != nil {
		h.Logger.Errorf("%v error bind json: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrGetRequest())
	}

	res, err := h.Service.GetList(ctx, input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, response_mapper.RenderStruct(http.StatusOK, res))
}
