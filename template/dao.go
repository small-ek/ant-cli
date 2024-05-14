package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// table 表名称
// GenDao 生成gorm
func GenDao(table string, tableStructure []TableStructure) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	preload := ""
	preloadImport := ""
	whereStr := ""
	for _, col := range tableStructure {
		if col.FieldType == "join" {
			preload = "Preload(clause.Associations)."
			preloadImport = `"gorm.io/gorm/clause"`
		}
		if col.IsSearch == 1 && col.Conditions != "" {
			whereStr += `sql.Where("` + col.FieldName + `", "` + col.Conditions + `", ` + table + `.` + utils.ToCamelCase(col.FieldName) + "),"
		}
	}

	return `package dao

import (
	"github.com/small-ek/antgo/db/adb/sql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"` + getFileName + `/app/model"
	` + preloadImport + `
)

type ` + humpTable + `Dao struct {
	db    *gorm.DB
	model *model.` + humpTable + `
}

func New` + humpTable + `Dao() *` + humpTable + `Dao {
	return &` + humpTable + `Dao{db: ant.Db()}
}

// Create
func (dao *` + humpTable + `Dao) Create(` + table + ` *model.` + utils.ToCamelCase(table) + `) error {
	return dao.db.Create(&` + table + `).Error
}

// DeleteById
func (dao *` + humpTable + `Dao) DeleteById(id int) error {
	return dao.db.Delete(&dao.model, id).Error
}

// Update
func (dao *` + humpTable + `Dao) Update(` + table + ` model.` + humpTable + `) error {
	return dao.db.Updates(&` + table + `).Error
}

// GetList
func (dao *` + humpTable + `Dao) GetList() (list []model.` + humpTable + `) {
	dao.db.Model(&dao.model).Find(&list)
	return list
}

// GetPage
func (dao *` + humpTable + `Dao) GetPage(page page.PageParam, ` + table + ` model.` + humpTable + `) (list []model.` + humpTable + `, total int64, err error) {
	err = dao.db.Model(&dao.model).Scopes(
		` + whereStr + `
		sql.Filters(page.Filter),
		sql.Order(page.Order),
		sql.Paginate(page.PageSize, page.CurrentPage),
	).` + preload + `Find(&list).Offset(0).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *` + humpTable + `Dao) GetById(id int) (row model.` + humpTable + `) {
	dao.db.Model(&dao.model).Where("id=?", id).Limit(1).` + preload + `Find(&row)
	return row
}
`
}
