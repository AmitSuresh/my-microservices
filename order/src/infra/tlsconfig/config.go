package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"go.uber.org/zap"
)

func LoadTLSConfig(l *zap.Logger, cfg *TLSConfig) (*tls.Config, error) {
	serverCert, err := tls.LoadX509KeyPair(cfg.ServerCertFile, cfg.ServerKeyFile)
	if err != nil {
		l.Fatal("Error loading server certificates", zap.Error(err))
		return nil, err
	}

	caCert, err := os.ReadFile(cfg.CACertFile)
	if err != nil {
		l.Fatal("Error loading CA cert", zap.Error(err))
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		l.Fatal("Failed to append CA certs")
		return nil, err
	}

	tlsConfig := &tls.Config{
		ClientCAs:          caCertPool,                     // Set the CA pool for client cert validation
		ClientAuth:         tls.RequireAndVerifyClientCert, // Require and verify client certificate
		Certificates:       []tls.Certificate{serverCert},  // Use the server's certificate
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	}

	return tlsConfig, nil
}
