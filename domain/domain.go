package domain

import "errors"

func validate(paths [][]string) error {
	if len(paths) == 0 {
		return errors.New("the list of paths is empty")
	}
	for _, v := range paths {
		if len(v) == 0 {
			return errors.New("contains empty path")
		}
		if len(v) == 1 {
			return errors.New("contains a path with one airport")
		}
		if len(v) > 2 {
			return errors.New("contains a path with more than two airports")
		}
		if v[0] == v[1] {
			return errors.New("contains a path where source and destination are the same")
		}
	}
	return nil
}

func (d Domain) Sort(paths [][]string) ([]string, error) {
	err := validate(paths)
	if err != nil {
		return nil, err
	}
	ports := map[string]*port{}
	for _, v := range paths {
		_, ok := ports[v[0]]
		if !ok {
			ports[v[0]] = &port{
				hasSource: true,
			}
		}
		_, ok = ports[v[1]]
		if !ok {
			ports[v[1]] = &port{
				hasDestination: true,
			}
		}
		ports[v[0]].hasSource = true
		ports[v[1]].hasDestination = true
	}
	source := ""
	destination := ""
	for k, v := range ports {
		if v.hasSource && !v.hasDestination {
			source = k
		} else if !v.hasSource && v.hasDestination {
			destination = k
		}
	}

	if len(source) == 0 || len(destination) == 0 {
		return nil, errors.New("can not find the source and destination")
	}

	return []string{source, destination}, nil
}
