package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type Task struct {
	Title       string
	Done 	  bool
}

type Folder struct {
    Name  string
    Tasks []Task
}

func (t *Task) Complete() error{
    if t.Done {
        return errors.New("задача уже выполнена")
    }
    t.Done = true
    return nil
}

func createFolder(folders []Folder, name string) ([]Folder, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return folders, errors.New("название папки не может быть пустым")
    }
    folders = append(folders, Folder{Name: name, Tasks: []Task{}})
    return folders, nil
}

func deleteFolder(folders []Folder, index int) ([]Folder, error) {
    if index < 0 || index >= len(folders) {
        return folders, errors.New("неверный индекс папки")
    }
    return append(folders[:index], folders[index+1:]...), nil
}

func renameFolder(folders []Folder, index int, newName string) ([]Folder, error) {
    newName = strings.TrimSpace(newName)
    if index < 0 || index >= len(folders) {
        return folders, errors.New("неверный индекс папки")
    }
    if newName == "" {
        return folders, errors.New("название папки не может быть пустым")
    }
    folders[index].Name = newName
    return folders, nil
}

func createTask(name string) (Task, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return Task{}, errors.New("название задачи не может быть пустым")
    }
    return Task{Title: name}, nil
}

func ScanNumber(scanner *bufio.Scanner) (int, error) {
    if !scanner.Scan() {
        if err := scanner.Err(); err != nil {
            return 0, err
        }
        return 0, io.EOF
    }

    text := strings.TrimSpace(scanner.Text())
    number, err := strconv.Atoi(text)
    if err != nil {
        return 0, errors.New("ошибка при вводе числа")
    }
    return number, nil
}

func main() {
    folders := []Folder{
        {Name: "Входящие", Tasks: []Task{}},
        {Name: "Работа", Tasks: []Task{}},
        {Name: "Личное", Tasks: []Task{}},
    }
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n=== ПАПКИ ===")

        for i, folder := range folders {
            fmt.Printf("%v. %v\n", i+1, folder.Name)
        }
        fmt.Println("0. Выход")

        fmt.Print("\nВыбери папку: \n")

        choice, err := ScanNumber(scanner)
        if errors.Is(err, io.EOF) {
            fmt.Println("Ввод завершён")
            return
        }
        if err != nil {
            fmt.Println("Ошибка при вводе числа:", err)
            continue
        }

        if choice > 0 && choice <= len(folders) {
            fmt.Printf("\n[%v]\n", folders[choice-1].Name)
            for i, task := range folders[choice-1].Tasks {
                fmt.Printf("Задача №%v: %s - выполнена: %v\n", i+1, task.Title, task.Done)
            }
        folderMenu:
            for {
                fmt.Printf("1. Создать папку\n")
                fmt.Printf("2. Удалить папку\n")
                fmt.Printf("3. Переименовать папку\n")
                fmt.Printf("4. Создать задачу\n")
                fmt.Printf("5. Выполнить задачу\n")
                fmt.Printf("6. Назад\n")
                fmt.Printf("0. Выход\n")

                folderChoice, err := ScanNumber(scanner)
                if errors.Is(err, io.EOF) {
                    fmt.Println("Ввод завершён")
                    return
                }
                if err != nil {
                    fmt.Println("Ошибка при вводе числа:", err)
                    continue
                }

                switch folderChoice {
                case 1:
                    fmt.Println("\nВведите название новой папки: ")
                    if !scanner.Scan() {
                        if err := scanner.Err(); err != nil {
                            fmt.Println("Ошибка чтения:", err)
                        } else {
                            fmt.Println("Ввод завершён")
                        }
                        return
                    }
                    name := scanner.Text()

                    res, err := createFolder(folders, name)

                    if err != nil {
                        fmt.Println("Ошибка при создании папки:", err)
                    } else {
                        folders = res
                        fmt.Printf("Папка %v успешно создана\n", strings.TrimSpace(name))
                    }
                    break folderMenu

                case 2:
                    fmt.Println("\nВведите номер папки, чтобы удалить: ")
                    numDelete, err := ScanNumber(scanner)
                    if errors.Is(err, io.EOF) {
                        fmt.Println("Ввод завершён")
                      return
                    }
                    if err != nil {
                        fmt.Println("Ошибка при вводе числа:", err)
                        continue
                    }
                    if numDelete < 1 || numDelete > len(folders) {
                        fmt.Println("Неверный номер папки")
                        break folderMenu
                    }
                    numOld := folders[numDelete - 1].Name

                    res, err := deleteFolder(folders, numDelete-1)

                    if err != nil {
                        fmt.Println("Ошибка при удалении папки:", err)
                    } else {
                        folders = res
                        fmt.Printf("Папка %v успешно удалена\n", numOld)
                    }
                    break folderMenu

                case 3:
                    fmt.Println("\nВведите номер папки, чтобы переименовать: ")
                    numRename, err := ScanNumber(scanner)
                    if errors.Is(err, io.EOF) {
                        fmt.Println("Ввод завершён")
                        return
                    }
                    if err != nil {
                        fmt.Println("Ошибка при вводе числа:", err)
                        continue
                    }
                    if numRename < 1 || numRename > len(folders) {
                        fmt.Println("Неверный номер папки")
                        break folderMenu
                    }

                    fmt.Println("\nВведите новое название папки: ")
                    if !scanner.Scan() {
                        if err := scanner.Err(); err != nil {
                            fmt.Println("Ошибка чтения:", err)
                        } else {
                            fmt.Println("Ввод завершён")
                        }
                        return
                    }
                    newName := scanner.Text()

                    oldName := folders[numRename-1].Name
                    res, err := renameFolder(folders, numRename-1, newName)
                    if err != nil {
                        fmt.Println("Ошибка при переименовании папки:", err)
                    } else {
                        folders = res
                        fmt.Printf("Папка %v успешно переименована в %v\n", oldName, strings.TrimSpace(newName))
                    }
                    break folderMenu

                case 4:
                    fmt.Println("\nВведите название новой задачи: ")
                    if !scanner.Scan() {
                        if err := scanner.Err(); err != nil {
                            fmt.Println("Ошибка чтения:", err)
                        } else {
                            fmt.Println("Ввод завершён")
                        }
                        return
                    }
                    taskName := scanner.Text()

                    res, err := createTask(taskName)
                    if err != nil {
                        fmt.Println("Ошибка при создании задачи:", err)
                    } else {
                        folders[choice-1].Tasks = append(folders[choice-1].Tasks, res)
                        fmt.Printf("\nЗадача %v успешно создана в папке %v\n", strings.TrimSpace(taskName), folders[choice-1].Name)
                    }
                    break folderMenu
                
                case 5:
                    fmt.Println("\nВведите номер задачи, которую хотите выполнить: ")
                    taskIndex, err := ScanNumber(scanner)
                    if errors.Is(err, io.EOF) {
                        fmt.Println("Ввод завершён")
                        return
                    }
                    if err != nil {
                        fmt.Println("Ошибка при вводе числа:", err)
                        continue
                    }
                    if taskIndex < 1 || taskIndex > len(folders[choice-1].Tasks) {
                        fmt.Println("Неверный номер задачи")
                        break folderMenu
                    }
                    if err := folders[choice-1].Tasks[taskIndex-1].Complete(); err != nil {
                        fmt.Println("Ошибка при выполнении задачи:", err)
                    } else {
                        fmt.Printf("Задача '%s' успешно выполнена\n", folders[choice-1].Tasks[taskIndex-1].Title)
                    }
                    break folderMenu
                case 6:
                    fmt.Println("\nВозврат в главное меню")
                    break folderMenu

                case 0:
                    fmt.Println("\nВыход")
                    return

                default:
                    fmt.Println("Неизвестная команда")
                    break folderMenu
                }
            }

        } else if choice == 0 {
            fmt.Println("\nВыход")
            return

        } else {
            fmt.Println("\nТакой папки не существует")
            continue
        }
    }
}