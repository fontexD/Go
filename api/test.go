for _, demo2 := range demo {

		demo3 := []Templatedata{
			Templatedata{
				Name:   demo2.Name,
				Group:  demo2.Group,
				Env:    demo2.Env,
				Status: RedisHealth(demo2.Host),
			},
		}
		encodeJson, err := json.Marshal(demo3)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(encodeJson))
	}

}
