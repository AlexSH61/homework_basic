package api

import (
	"io"
	"net/http"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/handlers"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type APIServer struct {
	config   *Config
	logger   *zap.SugaredLogger
	router   *mux.Router
	store    *store.Store //db
	users    *handlers.UserHandler
	orders   *handlers.OrderHandler
	products *handlers.ProductHandler
}

func New(config *Config, userHandler *handlers.UserHandler, orderHandler *handlers.OrderHandler, productHandler *handlers.ProductHandler, log *zap.SugaredLogger) (*APIServer, error) {
	// logger, err := logger.NewLogger()
	// if err != nil {
	// 	return nil, err
	// }
	return &APIServer{
		config:   config,
		logger:   log,
		router:   mux.NewRouter(),
		store:    nil, // Store будет инициализирован позже
		users:    userHandler,
		orders:   orderHandler,
		products: productHandler,
	}, nil
}
func (s *APIServer) Start() error {
	s.configRouter()
	//config store
	if err := s.configStore(); err != nil {
		return err
	}

	if err := s.configLogger(); err != nil {
		return err
	}
	s.logger.Info("Starting server", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func NewLogger(config *Config) (*zap.SugaredLogger, error) {
	logLevel, err := zap.ParseAtomicLevel(config.LoggerLevel)
	if err != nil {
		return nil, err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = logLevel
	cfg.Encoding = "json"
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

func (s *APIServer) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/orders", s.orders.InsertOrder).Methods(http.MethodPost)
	s.router.HandleFunc("/order/{user_ID:[0-9]+}", s.orders.GetOrdersByUserID)
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *APIServer) configStore() error {
	st := store.New(s.config.Store, s.logger)
	//connect
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
