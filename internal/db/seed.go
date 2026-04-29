package db

import (
	"context"
	"log"
	"math/rand"
	"strconv"

	"github.com/edwinboon/gopher-social-api/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy", "mallory", "oscar", "peggy", "trent",
	"victor", "walter", "sybil", "carol", "dan", "erin", "faythe", "ginny", "harry", "isabel", "jack", "karen", "larry", "mike", "nancy", "olivia",
	"pat", "quinn", "ruth", "sam", "ted", "uma", "wendy", "xavier", "yvonne", "zach", "amber", "bruce", "claire", "dylan", "elena", "felix", "george",
	"hannah", "irene", "jason", "kyle",
}

var titles = []string{
	"Why I Stopped Multitasking (And You Might Too)",
	"10 Tiny Habits That Instantly Improve Your Day",
	"The Art of Saying No Without Feeling Guilty",
	"What I Learned From Walking Every Day for 30 Days",
	"Minimalism for Beginners: Less Stuff, More Calm",
	"How to Build a Morning Routine You’ll Actually Keep",
	"My Favorite Tools for Staying Organized",
	"How to Find Focus in a World Full of Notifications",
	"Budgeting Without Spreadsheet Stress: My Simple Method",
	"The Best Books I’ve Read This Year (So Far)",
	"Meal Prep for Normal People: Simple and Fast",
	"Why Your Calendar Shouldn’t Be Your Boss",
	"From Procrastination to Action: 5 Practical Steps",
	"What Nobody Tells You About Productivity",
	"A Weekend Offline: Here’s What Happened",
	"How to Set Goals You’ll Actually Achieve",
	"My Writing Workflow to Beat Writer’s Block",
	"Starting Running: My Honest Beginner Experience",
	"The Power of Boredom: Making Space for Better Ideas",
	"Which Apps I Deleted (And Why)",
}

var contents = []string{
	"I used to juggle tasks all day and still feel behind. Here’s what changed when I switched to single-tasking.",
	"These small habits take under five minutes, but they add up fast. Pick one and try it today.",
	"Saying no is a skill, not a personality trait. This is how I do it politely and confidently.",
	"Thirty days of walking sounded too simple to matter—until I felt the results. Here’s what surprised me most.",
	"Minimalism isn’t about owning nothing. It’s about keeping what supports your life and letting the rest go.",
	"A good morning routine should fit your real life. Here’s a simple structure you can adapt in 10 minutes.",
	"Organization isn’t about more apps—it’s about fewer, better systems. These are the tools I actually rely on.",
	"Notifications steal attention in tiny pieces. Here are practical ways to protect your focus and get more done.",
	"Budgeting doesn’t need to be complicated. This is the lightweight approach that helped me stay consistent.",
	"I kept track of the books that genuinely stood out this year. Here are the highlights and why they mattered.",
	"Meal prep can be easy, flexible, and low-effort. This is my go-to plan for busy weeks.",
	"Your calendar should support your priorities, not dictate them. Here’s how I take control of my schedule.",
	"Procrastination is usually a signal, not laziness. These steps help me start when I don’t feel like it.",
	"Most productivity advice is louder than it is helpful. Here’s what’s actually made a difference for me.",
	"Being offline felt uncomfortable at first, then refreshing. Here’s what I noticed after two days unplugged.",
	"Goals fail when they’re vague or unrealistic. This is how I set goals that stay motivating and measurable.",
	"Writer’s block isn’t always a lack of ideas—it’s often friction. Here’s the workflow that keeps me writing.",
	"Running as a beginner is awkward, but doable. This is what helped me stay injury-free and consistent.",
	"Boredom can be a creative advantage. Here’s how I create space for thinking without forcing it.",
	"Deleting apps gave me time back immediately. Here’s what I removed, what I kept, and what I learned.",
}

var tags = []string{
	"productivity", "self-improvement", "habits", "mindfulness", "mental-health", "focus", "time-management",
	"minimalism", "wellness", "fitness", "running", "walking", "routine", "goal-setting", "motivation", "writing", "creativity",
	"personal-finance", "budgeting", "technology",
}

var comments = []string{
	"This really resonated with me—thanks for putting it into words.",
	"I needed this reminder today. Simple, but powerful.",
	"Great post. I’m going to try this approach this week.",
	"I’ve been struggling with the same thing—glad I’m not alone.",
	"Love how practical this is. The examples helped a lot.",
	"Interesting perspective. I hadn’t thought about it like that.",
	"I tried something similar and it worked surprisingly well.",
	"This was a quick read but it gave me a lot to think about.",
	"Do you have any tips for staying consistent when life gets busy?",
	"I shared this with a friend—super helpful.",
	"This is exactly why I started simplifying my routine too.",
	"Your writing style is so clear and easy to follow.",
	"I disagree with one part, but overall it’s a solid take.",
	"Could you expand on how you measure progress with this?",
	"I’m bookmarking this to revisit later.",
	"The section about focus hit home for me.",
	"Thanks for being honest about what didn’t work as well.",
	"This feels realistic, not like the usual “hustle” advice.",
}

func Seed(store store.Store) {
	ctx := context.Background()

	users := generateUsers(50)

	// Insert users into the database
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Printf("Error inserting user %s: %v", user.Username, err)
			return
		}
	}

	posts := generatePosts(200, users)

	// Insert posts into the database
	for _, p := range posts {
		if err := store.Posts.Create(ctx, p); err != nil {
			log.Printf("Error inserting post %s: %v", p.Title, err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	// Insert comments into the database
	for _, c := range comments {
		if err := store.Comments.Create(ctx, c); err != nil {
			log.Printf("Error inserting comment for post %d: %v", c.PostID, err)
			return
		}
	}

	log.Printf("Successfully seeded database with %d users, %d posts, and %d comments", len(users), len(posts), len(comments))
}

func generateUsers(n int) []*store.User {
	users := make([]*store.User, n)

	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + strconv.Itoa(i/len(usernames)),
			Password: "password",
			Email:    usernames[i%len(usernames)] + strconv.Itoa(i/len(usernames)) + "@example.com",
		}
	}

	return users
}

func generatePosts(n int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, n)

	for i := 0; i < n; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				// for sake of simplicity, just using 2 random tags per post
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(n int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, n)

	for i := 0; i < n; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		cms[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: comments[rand.Intn(len(comments))],
		}

	}

	return cms
}
