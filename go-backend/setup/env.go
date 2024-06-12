package setup

type Env struct {
	TimeoutSeconds int
	PostsPerPage int
	DBpassword string
	TokenSecret string
	RunMigration bool
	TestMode bool
}

func NewEnv() *Env {
	return &Env{
		TimeoutSeconds: 2,
		PostsPerPage: 10,
		DBpassword: "0123456789sqa9876543210",
		TokenSecret: "secret",
		RunMigration: true,
		TestMode: true,
	}
}