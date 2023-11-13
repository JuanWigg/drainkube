package drain

type PDB struct {
	name        string
	matchLabels map[string]string
}

type Pod struct {
	name   string
	labels map[string]string
}

type Deploy struct {
	name        string
	matchLabels map[string]string
}

type Rollout struct {
	name        string
	matchLabels map[string]string
}

type HPA struct {
	name      string
	targetRef string
	kind      string
}
