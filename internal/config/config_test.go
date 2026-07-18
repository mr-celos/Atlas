package config

import (
	"testing"
)

func TestLoad_Default(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("PRODUCT_VERSION", "")

	cfg := Load()
	if cfg.Port != ":8080" {
		t.Errorf("Expected Port to be :8080, got %s", cfg.Port)
	}
	if cfg.Version != "Development" {
		t.Errorf("Expected Version to be Development, got %s", cfg.Version)
	}
}

func TestLoad_WithEnvVars(t *testing.T) {
	t.Setenv("PORT", ":1234")
	t.Setenv("PRODUCT_VERSION", "1.2.3")

	cfg := Load()
	if cfg.Port != ":1234" {
		t.Errorf("Expected Port to be :1234, got %s", cfg.Port)
	}
	if cfg.Version != "1.2.3" {
		t.Errorf("Expected Version to be 1.2.3, got %s", cfg.Version)
	}
}
