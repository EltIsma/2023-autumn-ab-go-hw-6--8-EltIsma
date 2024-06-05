package repositories_test

import (
	"homework/models"
	"homework/repositories"
	"homework/services"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)



type MyTestSuite struct {
	suite.Suite
	service repositories.Repository
}

func (suite *MyTestSuite) SetupSuite() {
	repo := repositories.NewDeviceService()
	suite.service = services.NewService(repo)
}

func (suite *MyTestSuite) TestCreate() {
	devices := []models.Device{
		{
			SerialNum: "123",
			Model:     "model1",
			IP:        "1.1.1.1",
		},
		{
			SerialNum: "124",
			Model:     "model2",
			IP:        "1.1.1.2",
		},
	}
	for _, d := range devices {
		err := suite.service.CreateDevice(d)
		if err != nil {
			suite.T().Errorf("unexpected error: %v", err)
		}
	}

}

func (suite *MyTestSuite) TestUpdate() {
	newDevice := models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.2",
	}
	err := suite.service.UpdateDevice(newDevice)
	if err != nil {
		suite.T().Errorf("unexpected error: %v", err)
	}

	gotDevice, err := suite.service.GetDevice(newDevice.SerialNum)
	if err != nil {
		suite.T().Errorf("unexpected error: %v", err)
	}

	if gotDevice != newDevice {
		suite.T().Errorf("new device %+#v not equal got device %+#v", newDevice, gotDevice)
	}
}

func (suite *MyTestSuite) TestDelete() {
	newDevice := models.Device{
		SerialNum: "124",
		Model:     "model2",
		IP:        "1.1.1.2",
	}

	err := suite.service.DeleteDevice(newDevice.SerialNum)
	if err != nil {
		suite.T().Errorf("unexpected error: %v", err)
	}

	_, err = suite.service.GetDevice(newDevice.SerialNum)
	if err == nil {
		suite.T().Error("want error, but got nil")
	}

}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}


func TestDeleteDeviceUnexisting(t *testing.T) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)

	err := service.DeleteDevice("123")
	if err == nil {
		t.Errorf("want error, but got nil")
	}
}

func TestUpdateDeviceUnexsting(t *testing.T) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	device := models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.1",
	}

	err := service.CreateDevice(device)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	newDevice := models.Device{
		SerialNum: "124",
		Model:     "model1",
		IP:        "1.1.1.2",
	}
	err = service.UpdateDevice(newDevice)
	if err == nil {
		t.Errorf("want err, but got nil")
	}
}

func TestCreateDuplicate(t *testing.T) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	wantDevice := models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.1",
	}

	err := service.CreateDevice(wantDevice)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = service.CreateDevice(wantDevice)
	if err == nil {
		t.Errorf("want error, but got nil")
	}

}
func FuzzGetDevice(f *testing.F) {
	repo := repositories.NewDeviceService()

	// Test for chicking ID.
	f.Fuzz(func(t *testing.T, seqNum string) {
		expect := models.Device{
			SerialNum: "124",
			Model:     "model2",
			IP:        "1.1.1.2",
		}

		err := repo.CreateDevice(expect)

		require.NoError(t, err)

		actual, err := repo.GetDevice(expect.SerialNum)

		require.Equal(t, actual, expect)
		require.NoError(t, err)

		err = repo.DeleteDevice(expect.SerialNum)

		require.NoError(t, err)

		actual, err = repo.GetDevice(expect.SerialNum)

		require.Nil(t, actual)
		require.Error(t, err)
	})
}


func BenchmarkCRUD(b *testing.B) {
	b.Run("Get device", BenchmarkGet)
	b.Run("Create device", BenchmarkCreate)
	b.Run("Update device", BenchmarkUpdate)
	b.Run("Delete device", BenchmarkDelete)
}

func BenchmarkGet(b *testing.B) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	wantDevice :=models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.1",
	}

	_ = service.CreateDevice(wantDevice)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = service.GetDevice(wantDevice.SerialNum)
	}
}

func BenchmarkCreate(b *testing.B) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	device :=models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.1",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = service.CreateDevice(device)
	}
}

func BenchmarkUpdate(b *testing.B) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	device := models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.1",
	}

	_ = service.CreateDevice(device)
	

	newDevice := models.Device{
		SerialNum: "123",
		Model:     "model2",
		IP:        "1.1.1.2",
	}

	_ = service.CreateDevice(device)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = service.UpdateDevice(newDevice)
	}
}

func BenchmarkDelete(b *testing.B) {
	repo := repositories.NewDeviceService()
	service := services.NewService(repo)
	device := models.Device{
		SerialNum: "123",
		Model:     "model1",
		IP:        "1.1.1.1",
	}

	_ = service.CreateDevice(device)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = service.DeleteDevice(device.SerialNum)
	}
}