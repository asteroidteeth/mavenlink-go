package mavenlink

type LookupResult struct {
	Key, Id string
}

type User struct {
	User_id, Full_name, Photo_path, Email_address, Headline string
}

type Users struct {
	Count   int
	Users   map[string]User
	Results []LookupResult
}

func (u *Users) GetUserIds() []string {
	var result []string = make([]string, 0, u.Count)
	for k := range u.Users {
		result = append(result, k)
	}
	return result
}

func (u *Users) GetUserNames() []string {
	var result []string = make([]string, 0, u.Count)
	for _, v := range u.Users {
		result = append(result, v.Full_name)
	}
	return result
}

type TimeEntry struct {
	Id                 string
	Created_at         string
	Updated_at         string
	Date_performed     string
	Story_id           string
	Time_in_minutes    int
	Billable           bool
	Notes              string
	Rate_in_cents      int
	Currency           string
	Currency_symbol    string
	Currency_base_unit int
	User_can_edit      bool
	Workspace_id       string
	User_id            string
}

type TimeEntries struct {
	Count        int
	Time_entries map[string]TimeEntry
	Results      []LookupResult
}

func (e *TimeEntries) EntriesForUser(u *User) []*TimeEntry {
	var result []*TimeEntry
	for _, v := range e.Time_entries {
		if v.User_id == u.User_id {
			result = append(result, &v)
		}
	}
	return result
}
