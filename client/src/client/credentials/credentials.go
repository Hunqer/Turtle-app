package credentials

type CredentialsHolder interface{
	GetName()string
	GetCurrentServer()(string, error)
}