package main

import "fmt"

const (
	// IsAdmin is a role that can do anything
	IsAdmin = 1 << iota // 1
	// IsModerator is a role that can perform certain tasks
	IsModerator //10
	// IsMember is a role that can only add posts
	IsMember //100
	// CanArchive role can archive
	CanArchive //1000
	// CanDelete role can delete
	CanDelete //10000
	// CanUnpublished role can unpublish
	CanUnpublished //100000
	// CanAdd a role that can add posts
)

// AddPost adds a post
func AddPost(roles byte) {
	if IsMember&roles == IsMember {
		fmt.Println("You are allowed to add posts")
	} else {
		fmt.Println("Not allowed")
	}
}

// ArchivePost archives a post
func ArchivePost(roles byte) {
	if CanArchive&roles == CanArchive || IsAdmin&roles == IsAdmin || IsModerator&roles == IsModerator {
		fmt.Println("You are allowed to archive post")
	} else {
		fmt.Println("Not allowed")
	}
}

// UnpublishPost unpublishes a post
func UnpublishPost(roles byte) {
	if CanUnpublished&roles == CanUnpublished || IsAdmin&roles == IsAdmin || IsModerator&roles == IsModerator {
		fmt.Println("You are allowed to unpublish post")
	} else {
		fmt.Println("Not allowed")
	}
}

// DeletePost deletes a post
func DeletePost(roles byte) {
	if CanDelete&roles == CanDelete || IsAdmin&roles == IsAdmin {
		fmt.Println("You are allowed to delete post")
	} else {
		fmt.Println("Not allowed")
	}
}
