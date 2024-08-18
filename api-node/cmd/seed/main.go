package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)

var (
	genres = []string{
		"Fashion", "Technology", "Sports", "Beauty", "Lifestyle",
		"Food", "Travel", "Fitness", "Music", "Art",
	}

	tags = map[string][]string{
		"Fashion":    {"clothing", "accessories", "trends", "style", "designer"},
		"Technology": {"gadgets", "software", "innovation", "AI", "robotics"},
		"Sports":     {"football", "basketball", "tennis", "athletics", "swimming"},
		"Beauty":     {"skincare", "makeup", "haircare", "cosmetics", "wellness"},
		"Lifestyle":  {"home decor", "relationships", "self-improvement", "productivity", "mindfulness"},
		"Food":       {"recipes", "restaurants", "cooking", "nutrition", "baking"},
		"Travel":     {"destinations", "adventure", "culture", "hotels", "budget travel"},
		"Fitness":    {"workout", "nutrition", "yoga", "running", "weightlifting"},
		"Music":      {"genres", "artists", "concerts", "instruments", "production"},
		"Art":        {"painting", "sculpture", "photography", "digital art", "exhibitions"},
	}
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	err = db.InitDB(
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Seed data
	seedUsers(100)
	seedCreators()
	seedCategories()
	seedArticles()
	seedShopsAndProducts()

	fmt.Println("Seeding completed successfully!")
}

func seedUsers(count int) {
	for i := 0; i < count; i++ {
		user := models.UserProfile{
			UserID:    faker.UUIDDigit(),
			Username:  faker.Username(),
			Bio:       faker.Sentence(),
			AvatarURL: faker.URL(),
		}
		result := db.DB.Create(&user)
		if result.Error != nil {
			log.Printf("Error creating user: %v", result.Error)
		}
	}
	log.Printf("Created %d users", count)
}

func seedCreators() {
	var users []models.UserProfile
	db.DB.Limit(10).Find(&users)

	for _, user := range users {
		creator := models.Creator{
			UserProfileID: user.ID,
		}
		result := db.DB.Create(&creator)
		if result.Error != nil {
			log.Printf("Error creating creator: %v", result.Error)
			continue
		}
	}
	log.Printf("Created 10 creators")
}

func seedCategories() {
	categories := []string{
		"Fashion", "Technology", "Sports", "Beauty", "Lifestyle",
		"Food", "Travel", "Fitness", "Music", "Art",
	}

	for _, category := range categories {
		result := db.DB.Create(&models.Category{Name: category})
		if result.Error != nil {
			log.Printf("Error creating category: %v", result.Error)
		}
	}
	log.Printf("Created categories")
}

func seedArticles() {
	var creators []models.Creator
	db.DB.Preload("UserProfile").Find(&creators)

	for _, creator := range creators {
		for i := 0; i < 2; i++ {
			content := generateArticleContent(genres[i])
			article := models.Article{
				CreatorID:  creator.ID,
				Title:      fmt.Sprintf("%s Article %d by %s", genres[i], i+1, creator.UserProfile.Username),
				Content:    content,
				Views:      rand.Intn(10000),
			}
			result := db.DB.Create(&article)
			if result.Error != nil {
				log.Printf("Error creating article: %v", result.Error)
				continue
			}

			// Add tags to the article
			genreTags := tags[genres[i]]
			for j := 0; j < 3; j++ {
				tag := models.Tag{Name: genreTags[rand.Intn(len(genreTags))]}
				db.DB.FirstOrCreate(&tag, tag)
				db.DB.Model(&article).Association("Tags").Append(&tag)
			}
		}
	}
	log.Printf("Created 20 articles (2 for each creator)")
}

func seedShopsAndProducts() {
	var creators []models.Creator
	db.DB.Preload("UserProfile").Find(&creators)

	for i, creator := range creators {
		// Create a shop for each creator
		shop := models.Shop{
			CreatorID:   creator.ID,
			Name:        fmt.Sprintf("%s's %s Shop", creator.UserProfile.Username, genres[i]),
			Description: fmt.Sprintf("Welcome to %s's shop featuring %s products!", creator.UserProfile.Username, genres[i]),
		}
		result := db.DB.Create(&shop)
		if result.Error != nil {
			log.Printf("Error creating shop: %v", result.Error)
			continue
		}

		// Create 5-10 products for each shop
		numProducts := rand.Intn(6) + 5
		for j := 0; j < numProducts; j++ {
			product := generateProduct(genres[i], shop.ID)
			result := db.DB.Create(&product)
			if result.Error != nil {
				log.Printf("Error creating product: %v", result.Error)
			}
		}
	}
	log.Printf("Created shops and products for all creators")
}

func generateProduct(genre string, shopID uint) models.Product {
	var category models.Category
	db.DB.Where("name = ?", genre).First(&category)

	return models.Product{
		ShopID:      shopID,
		CategoryID:  category.ID,
		Title:       fmt.Sprintf("%s Item", genre),
		Description: fmt.Sprintf("An amazing %s product that you'll love. Perfect for enthusiasts and beginners alike.", genre),
		Price:       float64(rand.Intn(10000)) / 100,
		ImageURL:    faker.URL(),
	}
}

func generateArticleContent(genre string) string {
	switch genre {
	case "Fashion":
		return "The world of fashion is ever-evolving, and this season is no exception. From bold prints to minimalist designs, designers are pushing boundaries and redefining style. One trend that's making waves is sustainable fashion. More and more brands are embracing eco-friendly materials and ethical production methods, proving that style and sustainability can go hand in hand. Another exciting development is the resurgence of vintage-inspired pieces, with a modern twist. Whether you're a trendsetter or prefer timeless classics, there's something for everyone in this season's diverse fashion landscape."

	case "Technology":
		return "In the rapidly advancing world of technology, artificial intelligence continues to be at the forefront of innovation. From smart homes to autonomous vehicles, AI is revolutionizing the way we live and work. One area seeing significant progress is natural language processing, with chatbots and virtual assistants becoming increasingly sophisticated. Meanwhile, the Internet of Things (IoT) is expanding, connecting more devices and creating smarter, more efficient systems. As we look to the future, emerging technologies like quantum computing and augmented reality promise to push the boundaries of what's possible even further."

	case "Sports":
		return "The world of sports is constantly evolving, with athletes pushing the limits of human performance. This year has seen remarkable achievements across various disciplines. In track and field, we've witnessed record-breaking sprints and jaw-dropping long jumps. The world of team sports has been equally exciting, with nail-biting finishes in basketball championships and soccer tournaments. Beyond the professional arena, there's been a growing emphasis on grassroots sports development, with initiatives aimed at encouraging youth participation and promoting healthy lifestyles. As we look ahead, the intersection of sports and technology promises to bring even more innovations to training methods and performance analysis."

	case "Beauty":
		return "The beauty industry is experiencing a revolution, with a focus on inclusivity and natural ingredients. Brands are expanding their shade ranges to cater to diverse skin tones, and there's a growing demand for gender-neutral products. Clean beauty continues to gain momentum, with consumers seeking products free from harmful chemicals. Skincare routines are becoming more sophisticated, with multi-step regimens and targeted treatments gaining popularity. In makeup, we're seeing a shift towards more natural, glowing looks, although bold, artistic expressions are also having their moment. The industry is also embracing technology, with AI-powered skin analysis and custom-blended products becoming increasingly common."

	case "Lifestyle":
		return "In today's fast-paced world, there's a growing movement towards mindful living and work-life balance. People are prioritizing self-care and mental health, incorporating practices like meditation and yoga into their daily routines. The concept of hygge, the Danish art of coziness, continues to influence home decor trends, with an emphasis on creating comfortable, inviting spaces. Sustainability is also at the forefront of lifestyle choices, from eco-friendly home products to reducing plastic use. In the culinary world, plant-based diets and locally sourced ingredients are gaining popularity. As we navigate the challenges of modern life, the focus is on creating meaningful experiences and fostering genuine connections."

	case "Food":
		return "The culinary world is experiencing a renaissance, with chefs and home cooks alike exploring new flavors and techniques. Farm-to-table dining continues to grow in popularity, emphasizing fresh, locally sourced ingredients. Fusion cuisine is pushing boundaries, blending traditional recipes with modern twists. Plant-based eating is no longer just a trend but a mainstream choice, with innovative meat alternatives and creative vegetable-centric dishes. Fermentation is having a moment, from kombucha to artisanal pickles, adding complex flavors and potential health benefits to our diets. In the world of beverages, craft cocktails and non-alcoholic 'mocktails' are elevating drink menus, featuring unique ingredients and stunning presentations."

	case "Travel":
		return "As the world reopens, travel is taking on new meanings and forms. Sustainable tourism is at the forefront, with travelers seeking eco-friendly accommodations and experiences that positively impact local communities. Off-the-beaten-path destinations are gaining popularity, as people look for unique, authentic experiences away from crowded tourist spots. Wellness travel is on the rise, combining traditional vacation activities with health-focused offerings like spa treatments, yoga retreats, and mindfulness workshops. Technology is also transforming the travel experience, from virtual reality previews of destinations to AI-powered personalized itineraries. As we explore the world, there's a growing appreciation for cultural immersion and responsible tourism practices."

	case "Fitness":
		return "The fitness industry is constantly evolving, with new trends and technologies shaping how we approach health and wellness. High-Intensity Interval Training (HIIT) remains popular, offering efficient, effective workouts for busy lifestyles. Wearable technology has become increasingly sophisticated, providing detailed insights into our health metrics and helping to optimize workouts. Virtual fitness classes and apps have made exercise more accessible than ever, allowing people to work out from the comfort of their homes. There's also a growing focus on holistic fitness, incorporating mental health and nutrition alongside physical exercise. As we look to the future, personalized fitness plans based on genetic data and AI analysis are set to revolutionize how we approach our health and fitness goals."

	case "Music":
		return "The music industry continues to evolve at a rapid pace, driven by technological advancements and changing consumer preferences. Streaming platforms have transformed how we discover and consume music, with algorithms introducing listeners to new artists and genres. Virtual concerts have gained traction, allowing fans to experience live performances from anywhere in the world. In terms of genres, we're seeing a beautiful fusion of styles, with artists blending elements from different musical traditions to create unique sounds. The DIY music scene is thriving, with independent artists leveraging social media and digital platforms to build their fan bases. As we look to the future, innovations like AI-composed music and immersive audio experiences promise to push the boundaries of what's possible in the world of music."

	case "Art":
		return "The art world is experiencing a digital revolution, with new mediums and platforms reshaping how art is created, shared, and experienced. Digital art and NFTs (Non-Fungible Tokens) have exploded onto the scene, challenging traditional notions of ownership and value in the art market. Augmented and virtual reality are opening up new possibilities for immersive art experiences, allowing viewers to interact with artworks in unprecedented ways. There's also a growing emphasis on inclusivity and diversity, with galleries and institutions showcasing a broader range of voices and perspectives. Street art continues to gain recognition, blurring the lines between high art and popular culture. As technology advances, we can expect to see even more innovative forms of artistic expression emerge, pushing the boundaries of creativity and challenging our perceptions of what art can be."

	default:
		return fmt.Sprintf("This is a fascinating article about %s. It explores the latest trends, innovations, and developments in this exciting field. From cutting-edge research to practical applications, this article covers it all. Whether you're a seasoned expert or a curious newcomer, you'll find valuable insights and thought-provoking ideas here. Stay tuned for more updates in the world of %s!", genre, genre)
	}
}