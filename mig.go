package tokenized_reinvestment

import "embed"

//go:embed migrations/*.sql
var EmbedMigrations embed.FS
