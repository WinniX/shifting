package controller

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/winnix/shifting/api/domain"
	infra "github.com/winnix/shifting/api/infrastructure"
	"github.com/winnix/shifting/api/models"
)

type createScheduleRequest struct {
	ValidFrom time.Time      `json:"valid_from" validate:"required,dateinfuture"`
	Shifts    []shiftRequest `json:"shifts" validate:"required,dive"`
}

type shiftRequest struct {
	StartAtHour   int `json:"start_at_hour" validate:"min=0,max=23"`
	StartAtMinute int `json:"start_at_minute" validate:"min=0,max=59"`
	EndAtHour     int `json:"end_at_hour" validate:"min=0,max=23"`
	EndAtMinute   int `json:"end_at_minute" validate:"min=0,max=59"`
	NOfPpl        int `json:"n_of_ppl" validate:"min=1"`
}

func GetWeekByDate(ctx *infra.Context) (*infra.HandlerResult, error) {
	strDate := ctx.GinCtx.Param("date")
	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		ctx.Logger.Sugar().Infof("GetWeekByDate: invalid request: %s\n", err.Error())
		return BadRequest("")
	}

	return Ok(date)
}

func GetShift(ctx *infra.Context) (*infra.HandlerResult, error) {
	strID := ctx.GinCtx.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		ctx.Logger.Sugar().Infof("GetShift: invalid request: %s\n", err.Error())
		return BadRequest("")
	}

	shift := domain.Shift{}
	if err := ctx.Db.First(&shift, id).Error; err != nil {
		ctx.Logger.Sugar().Errorf("can't get a shift: %s\n", err.Error())
		return ServerError()
	}

	return Ok(shift)
}

func CreateSchedule(ctx *infra.Context) (*infra.HandlerResult, error) {
	model := createScheduleRequest{}
	if err := ctx.GinCtx.BindJSON(&model); err != nil {
		ctx.Logger.Sugar().Infof("CreateSchedule: invalid request: %s\n", err.Error())
		return BadRequest("")
	}

	if err := ctx.Validator.Struct(model); err != nil {
		failRes := models.FailedResponse{Errors: make(map[string]interface{})}
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			failRes.Errors[e.Field()] = e.Tag()
		}

		ctx.Logger.Sugar().Infof("CreateSchedule: invalid request: %s\n", err.Error())
		return BadRequest("")
	}

	if model.ValidFrom.Weekday() != time.Monday {
		ctx.Logger.Sugar().Infof("CreateSchedule: new schedule should start on Monday\n")
		return BadRequest("Schedule should start on Monday")
	}

	if err := ctx.Db.Where("valid_from = ?", model.ValidFrom).Delete(&domain.Shift{}).Error; err != nil {
		ctx.Logger.Sugar().Errorf("failed deleting shifts: %s\n", err.Error())
		return ServerError()
	}

	var invalidateShifts []domain.Shift
	ctx.Db.Find(&invalidateShifts, "valid_until IS NULL")
	if len(invalidateShifts) > 0 {
		for i := range invalidateShifts {
			invalidateShifts[i].ValidUntil = sql.NullTime{Time: getDateOnly(model.ValidFrom), Valid: true}
		}
		if err := ctx.Db.Save(invalidateShifts).Error; err != nil {
			ctx.Logger.Sugar().Errorf("failed invalidating shifts: %s\n", err.Error())
			return ServerError()
		}
	}

	weeklyShifts := generateWeeklyShifts(model)
	if err := ctx.Db.Create(&weeklyShifts).Error; err != nil {
		ctx.Logger.Sugar().Errorf("failed creating shifts: %s\n", err.Error())
		return ServerError()
	}

	return Ok(weeklyShifts)
}

func generateWeeklyShifts(r createScheduleRequest) []domain.Shift {
	nullTime := sql.NullTime{
		Valid: false,
	}
	result := make([]domain.Shift, 0, 7*len(r.Shifts))
	for weekDay := 0; weekDay < 7; weekDay++ {
		for _, item := range r.Shifts {
			s := domain.Shift{
				StartsAt:    getTimeOnly(item.StartAtHour, item.StartAtMinute),
				EndsAt:      getTimeOnly(item.EndAtHour, item.EndAtMinute),
				ValidFrom:   getDateOnly(r.ValidFrom),
				ValidUntil:  nullTime,
				WeekDay:     time.Weekday(weekDay),
				RequiresPpl: item.NOfPpl,
				AccountCode: "AAAA",
				PropertyId:  "MUC",
			}
			result = append(result, s)
		}
	}

	return result
}

func getDateOnly(dt time.Time) time.Time {
	return time.Date(
		dt.Year(), dt.Month(), dt.Day(),
		0, 0,
		0, 0, time.UTC)
}

func getTimeOnly(h int, m int) time.Time {
	t := time.Time{}
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		h, m,
		0, 0, time.UTC)
}
