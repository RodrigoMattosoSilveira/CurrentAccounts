package utilities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	k "github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
)

func GetSession(c *fiber.Ctx) *session.Session {
	return c.Locals("session").(*session.Session)
}

func WithSession(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		c.Locals("session", sess)
		c.Locals(k.PERSON_ID, sess.Get(k.PERSON_ID))
		return c.Next()
	}
}