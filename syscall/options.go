package syscall

func Options() [EndOfSo]Option {
	var opts [EndOfSo]Option
	copy(opts[:], options[:])
	return opts
}

func OptionByKey(key int) Option {
	return options[key]
}
