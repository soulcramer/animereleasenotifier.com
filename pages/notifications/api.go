package notifications

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// CountUnseen sends the number of unseen notifications.
func CountUnseen(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	unseen := user.Notifications().CountUnseen()

	return ctx.Text(strconv.Itoa(unseen))
}

// MarkNotificationsAsSeen marks all notifications as seen.
func MarkNotificationsAsSeen(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	notifications := user.Notifications().Notifications()

	for _, notification := range notifications {
		notification.Seen = arn.DateTimeUTC()
		notification.Save()
	}

	return "ok"
}

// Latest returns the latest notifications.
func Latest(ctx *aero.Context) string {
	userID := ctx.Get("id")
	user, err := arn.GetUser(userID)

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid user ID")
	}

	notifications := user.Notifications().Notifications()

	// Sort by date
	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].Created > notifications[j].Created
	})

	if len(notifications) > maxNotifications {
		notifications = notifications[:maxNotifications]
	}

	return ctx.JSON(notifications)
}

// Test sends a test notification to the logged in user.
func Test(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	user.SendNotification(&arn.PushNotification{
		Title:   "Anime Notifier",
		Message: "Yay, it works!",
		Icon:    "https://" + ctx.App.Config.Domain + "/images/brand/220.png",
		Type:    arn.NotificationTypeTest,
	})

	return "ok"
}
