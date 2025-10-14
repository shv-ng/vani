package data

type Data struct {
	ServerName     string
	Version        string
	RepositoryLink string
}

func GetData() Data {
	return Data{
		ServerName:     "vani",
		Version:        "0.0.1",
		RepositoryLink: "https://github.com/shv-ng/vani",
	}
}
