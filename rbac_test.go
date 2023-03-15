package rbac

type UserServiceImpl struct {
	Writer Writer[User]
	Reader Reader[User]
}

func (u *UserServiceImpl) SetReader(reader Reader[User]) {
	u.Reader = reader
}

func (u *UserServiceImpl) SetWriter(writer Writer[User]) {
	u.Writer = writer
}

// TODO 实现一个满足接口需求的测试
