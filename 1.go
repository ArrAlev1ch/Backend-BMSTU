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

func createFolder(folders []string, name string) ([]string, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return folders, errors.New("название папки не может быть пустым")
    }
    folders = append(folders, name)
    return folders, nil
}

func deleteFolder(folders []string, index int) ([]string, error) {
    if index < 0 || index >= len(folders) {
        return folders, errors.New("неверный индекс папки")
    }
    return append(folders[:index], folders[index+1:]...), nil
}

func renameFolder(folders []string, index int, newName string) ([]string, error) {
    newName = strings.TrimSpace(newName)
    if index < 0 || index >= len(folders) {
        return folders, errors.New("неверный индекс папки")
    }
    if newName == "" {
        return folders, errors.New("название папки не может быть пустым")
    }
    folders[index] = newName
    return folders, nil
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
    folders := []string{"Входящие", "Работа", "Личное"}
    workIndex := 1
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n=== ПАПКИ ===")

        for i, folder := range folders {
            fmt.Printf("%v. %v\n", i+1, folder)
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
            fmt.Printf("\n[%v]\n", folders[choice-1])

            if choice-1 == workIndex {
                fmt.Println("задач пока нет - занятие 02")
            }
        folderMenu:
            for {
                fmt.Printf("1. Создать папку\n")
                fmt.Printf("2. Удалить папку\n")
                fmt.Printf("3. Переименовать папку\n")
                fmt.Printf("4. Назад\n")
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

                    res, err := deleteFolder(folders, numDelete-1)

                    if err != nil {
                        fmt.Println("Ошибка при удалении папки:", err)
                    } else {
                        folders = res

                        if numDelete-1 == workIndex {
                            workIndex = -1
                        } else if numDelete-1 < workIndex {
                            workIndex--
                        }

                        fmt.Printf("Папка %v успешно удалена\n", numDelete)
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

                    oldName := folders[numRename-1]
                    res, err := renameFolder(folders, numRename-1, newName)
                    if err != nil {
                        fmt.Println("Ошибка при переименовании папки:", err)
                    } else {
                        folders = res
                        fmt.Printf("Папка %v успешно переименована в %v\n", oldName, strings.TrimSpace(newName))
                    }
                    break folderMenu

                case 4:
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