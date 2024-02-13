package assets

import "embed"

//go:embed sql/migrations/*.sql
var Migrations embed.FS
