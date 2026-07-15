                  =============  PHASE ONE  ============
                 go run ./cmd/server
                         │
                         ▼
                    main.go starts
                         │
                         ▼
                 config.LoadConfig()
                         │
                         ▼
                   Read .env file
                         │
                         ▼
                  Config struct created
                         │
                         ▼
             database.Connect(cfg)
                         │
                         ▼
               PostgreSQL Connection
                         │
                         ▼
              router.SetupRouter()
                         │
                         ▼
                  Create Gin Engine
                         │
                         ▼
              Register /health Route
                         │
                         ▼
                Return Router Object
                         │
                         ▼
                 r.Run(":8080")
                         │
                         ▼
              Server starts listening
                         │
        ─────────────────┼─────────────────
                         │
             Browser sends GET /health
                         │
                         ▼
              Gin finds matching route
                         │
                         ▼
                 Handler executes
                         │
                         ▼
                  JSON Response sent
                         │
                         ▼
                    Browser displays


                    =============  PHASE 2 ===============

main.go

↓

Repository

↓

Service

↓

Handler

↓

Router