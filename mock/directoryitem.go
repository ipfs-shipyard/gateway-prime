package mock

type directoryItem struct {
	Size      string
	Name      string
	Path      string
	Hash      string
	ShortHash string
}

func (d *directoryItem) GetSize() string {
	return d.Size
}

func (d *directoryItem) GetName() string {
	return d.Name
}

func (d *directoryItem) GetPath() string {
	return d.Path
}

func (d *directoryItem) GetHash() string {
	return d.Hash
}

func (d *directoryItem) GetShortHash() string {
	return d.ShortHash
}
