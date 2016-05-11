package db

import (
	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"

	_ "github.com/Sirupsen/logrus"
)

func (maker *PersonMaker) Create(person *Person) error {
	query, _, err := sq.
		Insert("person").Columns("first_name", "last_name", "username", "password").Values("first_name", "last_name", "username", "password").
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, person.FirstName, person.LastName, person.Username, string(person.Password))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	person.ID = int(id)

	return nil
}

func (maker *PersonMaker) GetByID(id int) (*Person, error) {
	query, _, err := sq.
		Select("id", "first_name", "last_name", "username", "created_at", "last_updated").From("person").
		Where(sq.Eq{"ID": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	var person = &Person{}
	err = util.GetAndMarshal(query, maker, person, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (maker *CourseMaker) CreateCourse(course *Course) error {

	query, _, err := sq.
		Insert("course").Columns("name").Values("name").
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, course.Name)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	course.ID = int(id)

	return nil

}

func (maker *CourseMaker) UpdateCourse(course *Course) error {

	query, _, err := sq.
		Update("course").
		Set("name", course.Name).
		Where(sq.Eq{"id": course.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = util.PrepAndExec(query, maker, course.Name, course.ID)
	if err != nil {
		return err
	}

	return nil
}

func (maker *CourseMaker) GetCourseByID(id int) (*Course, error) {
	query, _, err := sq.
		Select("id, name, created_at, last_updated").From("course").
		Where(sq.Eq{"ID": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	var course = &Course{}
	err = util.GetAndMarshal(query, maker, course, id)
	if err != nil {
		return nil, err
	}

	return course, nil

}
func (maker *CourseMaker) DestroyCourse(course *Course) error {
	query, _, err := sq.
		Delete("course").
		Where(sq.Eq{"ID": course.ID}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = util.PrepAndExec(query, maker, course.ID)
	if err != nil {
		return err
	}

	return nil
}

func (maker *TeacherMaker) CreateTeacher(teacher *Teacher) error {
	return nil

}
func (maker *TeacherMaker) UpdateTeacher(teacher *Teacher) error {
	return nil

}
func (maker *TeacherMaker) GetTeacherByID(id int) (*Teacher, error) {
	return nil, nil

}

func (maker *TeacherMaker) DestroyTeacher(t *Teacher) error {
	return nil
}

func (maker *StudentMaker) CreateStudent(student *Student) error {
	return nil

}
func (maker *StudentMaker) UpdateStudent(student *Student) error {
	return nil

}
func (maker *StudentMaker) GetStudentByID(id int) (*Student, error) {
	return nil, nil

}
func (maker *StudentMaker) DestroyStudent(student *Student) error {
	return nil
}

type AssignmentMaker struct {
	*sqlx.DB
}

func (maker *AssignmentMaker) CreateAssignment(assig *Assignment) error {
	query, _, err := sq.
		Insert("assignment").Columns("student_id", "teacher_id").
		Values(assig.StudentID, assig.TeacherID).
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, assig.StudentID, assig.TeacherID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	assig.ID = int(id)

	return nil
}

func (maker *AssignmentMaker) UpdateAssignment(assig *Assignment) error {
	return nil

}

func (maker *AssignmentMaker) GetAssignmentByID(id int) (*Assignment, error) {
	return nil, nil
}

func (maker *AssignmentMaker) DestroyAssignment(assig *Assignment) error {
	return nil
}
