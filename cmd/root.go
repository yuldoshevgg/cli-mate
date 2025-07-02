package cmd

import (
	"cli-mate/internal"
	"cli-mate/internal/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var (
	version string = "2.1.0"
	Debug   bool
	rootCmd = cobra.Command{
		Use:     "cli-mate",
		Short:   "cli-mate - Simple weather application for your terminal",
		Version: version,
		PreRun:  toggleDebug,
		Example: "cli-mate",
		Run:     runRoot,
	}
)

func init() {
	cfg = config.Load()
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "verbose logging")
}

func runRoot(cmd *cobra.Command, args []string) {
	location, err := internal.GetGeoLocation()
	if err != nil {
		log.Fatalf("failed to get geo location: %v", err)
	}

	fmt.Printf("\n\n github_token: %v, api_key: %v, url: %v\n\n", cfg.GithubToken, cfg.APIKey, cfg.URL)

	airPollution, err := internal.AirPollution(cfg, location.Lat, location.Lon)
	if err != nil {
		log.Fatalf("failed to get air pollution: %v", err)
	}

	weather, err := internal.GetWeather(cfg, location.Lat, location.Lon)
	if err != nil {
		log.Fatalf("failed to get weather: %v", err)
	}

	fmt.Printf("📍 %v %v°C\n💨 %v m/s\n💧 %v%%\n🌅 Sunrise: %v\n🌆 Sunset: %v\n\nAir Pollution Index: %v %v%v\n",
		location.City,
		int32(weather.Main.Temp),
		weather.Wind.Speed,
		weather.Main.Humidity,
		time.Unix(int64(weather.Sys.Sunrise), 0).Format("15:04"),
		time.Unix(int64(weather.Sys.Sunset), 0).Format("15:04"),
		internal.AirPollutionMainColor[airPollution.List[0].Main.Aqi].Emoji,
		internal.AirPollutionMainColor[airPollution.List[0].Main.Aqi].Color,
		internal.AirPollutionMainColor[airPollution.List[0].Main.Aqi].Text,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
