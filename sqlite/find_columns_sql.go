package sqlite

var findColumnSql = `SELECT sql FROM sqlite_master WHERE type = 'table' AND tbl_name = ?`
