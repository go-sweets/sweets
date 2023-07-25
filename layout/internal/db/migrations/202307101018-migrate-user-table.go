package migrations

type UserMigration struct{}

func (m *UserMigration) Name() string {
	return "2023071010190000-migrate-user-table"
}

func (m *UserMigration) Up() error {
	return nil
}

func (m *UserMigration) Down() error {
	return nil
}
