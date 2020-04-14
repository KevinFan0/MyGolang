var mu sync.Mutex				// guards icons
var icons map[string]image.Image

// 保证所有goroutine能够观察到loadIcons效果的方式，是用一个mutex来同步检查。
func Icon(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

func loadIcons() {
    icons = make(map[string]image.Image)
    icons["spades.png"] = loadIcon("spades.png")
    icons["hearts.png"] = loadIcon("hearts.png")
    icons["diamonds.png"] = loadIcon("diamonds.png")
    icons["clubs.png"] = loadIcon("clubs.png")
}


// 引入一个允许多读的锁，使得可以对该变量进行并发访问
var mu sync.RWMutex
var icons map[string]image.Image
func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()
	// acquire an exclusive lock
	mu.Lock()
	if icons == nil {				// 重新检查icons是否为nil，防止其他go程在获取所之前初始化变量
		loadIcons()
	}
	icons := icons[name]
	mu.Unlock()
	return icon
}


var loadIconsOnce sync.Once
var icons map[string]image.Image
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}