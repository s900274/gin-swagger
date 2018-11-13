package services

import (
	"github.com/gin-gonic/gin"
	"httpservice/model"
	"database/sql"
	"fmt"
	"net/http"
	logger "github.com/shengkehua/xlog4go"
	"define"
)

type BuildingController struct {

}

func returnErrorResp(ctx *gin.Context, err error) {
	logger.Error("System Error : %v", err)
	ctx.JSON(http.StatusBadRequest, model.ApiResponse{
		define.ERR_ERROR,
		fmt.Sprintf("%v", err),
		"",
	})
}

// @Summary 列出所有樓層公司名稱
// @Produce  json
// @Tags 樓層公司資訊
// @Success 200 {object} model.ApiResponse
// @Failure 400
// @Router /api/building/company [get]
func (c *BuildingController)GetBuildingCompany(ctx *gin.Context) {

	// Open Sqlite
	db, err := sql.Open("sqlite3", "./db/building.db")

	// If error then return error message
	if err != nil {
		returnErrorResp(ctx, err)
		return
	}

	// Query Building Companies
	rows, err := db.Query("SELECT floor, name FROM companies")

	// If error then return error message
	if err != nil {
		returnErrorResp(ctx, err)
		return
	}

	// Close Sqlite
	db.Close()

	logger.Debug("Fetch result : %v", rows)

	// Append the result into slice
	arrCompanyInfo := make([]model.CompanyInfo, 0)
	for rows.Next() {
		var floor string
		var name string
		rows.Scan(&floor, &name)

		companyInfo := model.CompanyInfo{}
		companyInfo.Floor = floor
		companyInfo.Name  = name

		arrCompanyInfo = append(arrCompanyInfo, companyInfo)
	}

	// Return result
	ctx.JSON(200, model.ApiResponse{
		0,
		"SUCCESS",
		arrCompanyInfo,
	})
	return
}

// @Summary 新增公司
// @Accept  json
// @Produce  json
// @Tags 樓層公司資訊
// @Param   sendBody body model.CompanyInfo true  "body"
// @Success 200 {object} model.ApiResponse
// @Failure 400
// @Router /api/building/company [post]
func (c *BuildingController)AddBuildingCompany(ctx *gin.Context) {

	// Verify the request struct is validate
	var req = &model.CompanyInfo{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		returnErrorResp(ctx, err)
		return
	}

	// Open Sqlite
	db, err := sql.Open("sqlite3", "./db/building.db")
	if err != nil {
		returnErrorResp(ctx, err)
		return
	}

	// Insert the data into table
	stmt, _ := db.Prepare("INSERT INTO companies(floor, name) values(?,?)")

	_, err = stmt.Exec(req.Floor, req.Name)
	if err != nil {
		returnErrorResp(ctx, err)
		return
	}

	// Close Sqlite
	db.Close()

	// Return result
	ctx.JSON(http.StatusBadRequest, model.ApiResponse{
		define.ERR_OK,
		"SUCCESS",
		"",
	})
}