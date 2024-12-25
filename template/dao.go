package template

import (
	"github.com/small-ek/ant-cli/utils"
)

// table 表名称
// GenDao 生成gorm
func GenDao(table string, tableStructure []TableStructure) string {
	getFileName := utils.GetFileName()
	humpTable := utils.ToCamelCase(table)
	smallHumpTable := utils.ToCamelCaseLower(table)
	preload := ""
	preloadImport := ""
	whereStr := ""
	for _, col := range tableStructure {
		if col.FieldType == "join" {
			preload = "Preload(clause.Associations)."
			preloadImport = `"gorm.io/gorm/clause"`
		}
		if col.IsSearch == 1 && col.Conditions != "" {
			whereStr += `asql.Where("` + col.FieldName + `", "` + col.Conditions + `", page.FilterMap["` + col.FieldName + `"]` + `),`
		}
	}

	return `package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"` + getFileName + `/app/entity/models"
	` + preloadImport + `
)

type ` + humpTable + `Dao struct {
	db    *gorm.DB
	models *models.` + humpTable + `
}

func New` + humpTable + `Dao(db *gorm.DB) *` + humpTable + `Dao {
	if db == nil {
		db = ant.Db()
	}
	return &` + humpTable + `Dao{db: db}
}

// Create
func (dao *` + humpTable + `Dao) Create(` + smallHumpTable + ` models.` + utils.ToCamelCase(table) + `) error {
	return dao.db.Create(&` + smallHumpTable + `).Error
}

// DeleteById
func (dao *` + humpTable + `Dao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *` + humpTable + `Dao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *` + humpTable + `Dao) Update(` + smallHumpTable + ` models.` + humpTable + `) error {
	return dao.db.Updates(&` + smallHumpTable + `).Error
}

// GetList
func (dao *` + humpTable + `Dao) GetList() (list []models.` + humpTable + `) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *` + humpTable + `Dao) GetPage(page page.PageParam) (list []models.` + humpTable + `, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		` + whereStr + `
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).` + preload + `Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *` + humpTable + `Dao) GetById(id int) (row models.` + humpTable + `) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).` + preload + `Find(&row)
	return row
}
`
}
