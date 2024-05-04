package util

func GetTotalNumberOfTasks(tfs []TaskFile) int {
	total := 0
	for _, tf := range tfs {
		total += len(tf.Tasks)
	}
	return total
}
