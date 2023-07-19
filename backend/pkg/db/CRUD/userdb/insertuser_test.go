package userdb_test

const MIGRATIONS_FILE_PATH = "../../migrations"

const BIG_ABOUT_ME = `Hello there! I am Taylor, a widely traveled time-jaunting omniscient dream-weaver. When I am not summoning sentient storms on a distant alien planet or deciphering ancient prophecies in forgotten crypts, you'll find me at the edge of reality, watching the universe unfurl its tapestry of starlight and darkness.

Born and raised in the vibrant city of Neon Noir, I am a part-time gondolier in the phosphorescent canals that criss-cross our sprawling metropolis. Here, I guide countless souls beneath towering skyscrapers, their reflections dancing upon the undulating water as they tell me their life stories. These stories are the fuel that lights up my ever-insatiable curiosity and molds the foundations of my every adventure.

As a connoisseur of cosmic curiosities, I have amassed a collection of infinity stones, stardust from dying supernovae, and a peculiar talking bonsai from a hyper-advanced civilization that exists 100,000 light years away. I am also an aspiring multi-dimensional chef, specializing in preparing exotic, out-of-this-world dishes, like Pulsar Pasta and Nebula Nougat. My food can cause taste buds to witness the Big Bang and transcend the known sensory dimensions.

I spend my spare time crafting sonic symphonies from the hum of celestial bodies and the whispers of the cosmic winds. My compositions, they say, can soothe even the most tempestuous black holes and harmonize the discordant galaxies.

Through my endless adventures, I've adopted a quantum cat named Schrody. Schrody loves lounging on the edge of space-time, and her favorite toy is a pocket universe I've fashioned just for her amusement.

In the realm of academia, I hold a degree in Quantum Anthropology from the University of Intergalactic Studies, with a minor in Temporal Linguistics. Currently, I am pursuing my Ph.D. on 'The Sociocultural Impact of Wormhole Immigration on Interstellar Trade Dynamics'.

At heart, I am an interstellar storyteller, narrating the untold tales of the cosmos, unveiling the mysteries of existence, and creating a narrative that weaves through the very fabric of space and time. My mission is to promote understanding and harmony among the vast and diverse life-forms that make up our multifaceted, marvelous universe.

Oh, and I make a killer cup of starlight espresso, the perfect fuel for any cosmic voyage! Join me, and let's write the next chapter of our shared cosmic story together.`

// func TestInsertUser(t *testing.T) {
// 	tempDir := t.TempDir()
// 	tempDBFilePath := &helpermodels.FilePathComponents{
// 		Directory: tempDir,
// 		FileName:  "test",
// 		Extension: ".db",
// 	}
// 	err := dbutils.CreateDatabase(tempDBFilePath)
// 	if err != nil {
// 		t.Errorf("Failed to create test database: %s", err)
// 	}
// 	migrationConstructor := &dbmodels.NativeMigrate{}
// 	migrateUpDown := &dbmodels.NativeMigrateUpdates{}
// 	err = dbutils.MigrateChangesUp(tempDBFilePath, MIGRATIONS_FILE_PATH, migrationConstructor, migrateUpDown)
// 	if err != nil {
// 		t.Errorf("failed to migrate changes up to test database: %s", err)
// 	}
// 	testDB, err := sql.Open("sqlite3", tempDBFilePath.Directory+"/"+tempDBFilePath.FileName+tempDBFilePath.Extension)
// 	if err != nil {
// 		t.Errorf("Failed to open test database: %s", err)
// 	}
// 	testCases := []struct {
// 		name        string
// 		user        *dbmodels.User
// 		columnType  string
// 		searchValue string
// 		expectError bool
// 	}{
// 		{
// 			name: "Correct user data",
// 			user: &dbmodels.User{
// 				UserId:         "1",
// 				IsLoggedIn:     0,
// 				Email:          "user@test.com",
// 				HashedPassword: "hashed_password",
// 				FirstName:      "First",
// 				LastName:       "Last",
// 				DOB:            time.Now(),
// 				AvatarPath:     "path/to/avatar",
// 				DisplayName:    "User",
// 				AboutMe:        "About me",
// 			},
// 			columnType:  "user_id",
// 			searchValue: "1",
// 			expectError: false,
// 		},
// 		{
// 			name: "Correct user data2",
// 			user: &dbmodels.User{
// 				UserId:         "2",
// 				IsLoggedIn:     1,
// 				Email:          "user2@test.com",
// 				HashedPassword: "43987ohflkshfsdugyi8a:_dsfkjhdfsdfjshdkfah!dsjfhsdjkfhskjdfhksjdfhksdjfh",
// 				FirstName:      "First2",
// 				LastName:       "Last2",
// 				DOB:            time.Now(),
// 				AvatarPath:     "path/to/avatar2",
// 				DisplayName:    "User2",
// 				AboutMe:        BIG_ABOUT_ME,
// 			},
// 			columnType:  "about_me",
// 			searchValue: BIG_ABOUT_ME,
// 			expectError: false,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err = userdb.InsertUser(testDB, tc.user)
// 			if tc.expectError && err == nil {
// 				t.Error("Expected an error, but got nil")
// 			} else if !tc.expectError && err != nil {
// 				t.Errorf("Unexpected error: %s", err)
// 			}
// 			_, err := userdb.SelectUser(testDB, tc.columnType, tc.searchValue)
// 			if err != nil {
// 				t.Errorf("Unexpected error: %s", err)
// 			}
// 		})
// 	}
// }
