package webhook

import (
	"testing"

	"github.com/kubeflow/spark-operator/api/v1beta2"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAddAffinity(t *testing.T) {
	// Given
	testAffinity := corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{}}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: nil},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: nil},
			},
		},
	}

	// When
	patchOps := addAffinity(testSparkApp, testAffinity, true)

	// Then
	assert.Equal(t, 2, len(patchOps))
	assert.Equal(t, "add", patchOps[0].Op)
	assert.Equal(t, "/spec/driver/affinity", patchOps[0].Path)
	assert.Equal(t, testAffinity, patchOps[0].Value)
	assert.Equal(t, "add", patchOps[1].Op)
	assert.Equal(t, "/spec/executor/affinity", patchOps[1].Path)
	assert.Equal(t, testAffinity, patchOps[1].Value)
}

func TestAddAffinity_WithExist(t *testing.T) {
	// Given
	testAffinity := corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{}}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{PreferredDuringSchedulingIgnoredDuringExecution: []corev1.PreferredSchedulingTerm{}}}},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{PreferredDuringSchedulingIgnoredDuringExecution: []corev1.PreferredSchedulingTerm{}}}},
			},
		},
	}

	// When
	patchOps := addAffinity(testSparkApp, testAffinity, true)

	// Then
	assert.Equal(t, 2, len(patchOps))
	assert.Equal(t, "replace", patchOps[0].Op)
	assert.Equal(t, "/spec/driver/affinity", patchOps[0].Path)
	assert.Equal(t, testAffinity, patchOps[0].Value)
	assert.Equal(t, "replace", patchOps[1].Op)
	assert.Equal(t, "/spec/executor/affinity", patchOps[1].Path)
	assert.Equal(t, testAffinity, patchOps[1].Value)
}

func TestAddAffinity_Soft(t *testing.T) {
	// Given
	testAffinity := corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{}}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: nil},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: nil},
			},
		},
	}

	// When
	patchOps := addAffinity(testSparkApp, testAffinity, false)

	// Then
	assert.Equal(t, 2, len(patchOps))
	assert.Equal(t, "add", patchOps[0].Op)
	assert.Equal(t, "/spec/driver/affinity", patchOps[0].Path)
	assert.Equal(t, testAffinity, patchOps[0].Value)
	assert.Equal(t, "add", patchOps[1].Op)
	assert.Equal(t, "/spec/executor/affinity", patchOps[1].Path)
	assert.Equal(t, testAffinity, patchOps[1].Value)
}

func TestAddAffinity_Soft_WithExist(t *testing.T) {
	// Given
	testAffinity := corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{}}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{PreferredDuringSchedulingIgnoredDuringExecution: []corev1.PreferredSchedulingTerm{}}}},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{PreferredDuringSchedulingIgnoredDuringExecution: []corev1.PreferredSchedulingTerm{}}}},
			},
		},
	}

	// When
	patchOps := addAffinity(testSparkApp, testAffinity, false)

	// Then
	assert.Equal(t, 0, len(patchOps))
}

func TestAddToleration(t *testing.T) {
	// Given
	testToleration := []corev1.Toleration{{}, {}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: nil},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: nil},
			},
		},
	}

	// When
	patchOps := addToleration(testSparkApp, testToleration, true)

	// Then
	assert.Equal(t, 2, len(patchOps))
	assert.Equal(t, "add", patchOps[0].Op)
	assert.Equal(t, "/spec/driver/toleration", patchOps[0].Path)
	assert.Equal(t, testToleration, patchOps[0].Value)
	assert.Equal(t, "add", patchOps[1].Op)
	assert.Equal(t, "/spec/executor/toleration", patchOps[1].Path)
	assert.Equal(t, testToleration, patchOps[1].Value)
}

func TestAddToleration_WithExist(t *testing.T) {
	// Given
	testToleration := []corev1.Toleration{{}, {}}
	existToleration := []corev1.Toleration{{Key: "test1"}, {Key: "test2"}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: existToleration},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: existToleration},
			},
		},
	}

	// When
	patchOps := addToleration(testSparkApp, testToleration, true)

	// Then
	assert.Equal(t, 2, len(patchOps))
	assert.Equal(t, "replace", patchOps[0].Op)
	assert.Equal(t, "/spec/driver/toleration", patchOps[0].Path)
	assert.Equal(t, testToleration, patchOps[0].Value)
	assert.Equal(t, "replace", patchOps[1].Op)
	assert.Equal(t, "/spec/executor/toleration", patchOps[1].Path)
	assert.Equal(t, testToleration, patchOps[1].Value)
}

func TestAddToleration_Soft(t *testing.T) {
	// Given
	testToleration := []corev1.Toleration{{}, {}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: nil},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: nil},
			},
		},
	}

	// When
	patchOps := addToleration(testSparkApp, testToleration, false)

	// Then
	assert.Equal(t, 2, len(patchOps))
	assert.Equal(t, "add", patchOps[0].Op)
	assert.Equal(t, "/spec/driver/toleration", patchOps[0].Path)
	assert.Equal(t, testToleration, patchOps[0].Value)
	assert.Equal(t, "add", patchOps[1].Op)
	assert.Equal(t, "/spec/executor/toleration", patchOps[1].Path)
	assert.Equal(t, testToleration, patchOps[1].Value)
}

func TestAddToleration_Soft_WithExist(t *testing.T) {
	// Given
	testToleration := []corev1.Toleration{{}, {}}
	existToleration := []corev1.Toleration{{Key: "test1"}, {Key: "test2"}}
	testSparkApp := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: v1beta2.SparkApplicationSpec{
			Driver: v1beta2.DriverSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: existToleration},
			},
			Executor: v1beta2.ExecutorSpec{
				SparkPodSpec: v1beta2.SparkPodSpec{Tolerations: existToleration},
			},
		},
	}

	// When
	patchOps := addToleration(testSparkApp, testToleration, false)

	// Then
	assert.Equal(t, 0, len(patchOps))
}
