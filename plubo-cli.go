package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
	"github.com/fatih/color"
)

//go:embed functionalities/*.php components/*.php alpine/alpine.ts tailwind/*
var content embed.FS

func main() {
	args, err := script.Args().Slice()

	if err != nil {
		fmt.Println(err)
	}

	if len(args) > 1 {
		switch args[0] {
		case "create", "functionality", "func", "add":
			create(args[1])
		case "component", "comp":
			component(args[1])
		case "utils":
			utils(args[1])
		case "lib", "module":
			lib(args[1])
		}

	} else if len(args) > 0 {
		switch args[0] {
		case "init_template", "init-template", "init-repo", "init_repo", "namespace-project":
			init_template()
		}
	} else {
		color.Cyan("\nTo the rational mind, nothing is inexplicable; only unexplained.\n\n")
	}

}

func create(functionality string) {

	exPath, _ := script.Exec("pwd").Exec("tr -d '\n'").String()
	new_text := ""

	switch functionality {
	case "admin-menus", "adminmenu", "admin-menu", "menus":
		original_text, _ := content.ReadFile("functionalities/AdminMenus.php")
		file, err := os.Create(exPath + "/Functionality/AdminMenus.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("ADMIN MENUS CLASS CREATED")
		}
		file.Close()
	case "ajax", "ajax-actions", "ajax-action":
		original_text, _ := content.ReadFile("functionalities/AjaxActions.php")
		file, err := os.Create(exPath + "/Functionality/AjaxActions.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("AJAX ACTIONS CLASS CREATED")
		}
		file.Close()
	case "api", "endpoint", "endpoints", "api-endpoint", "api-endpoints":
		original_text, _ := content.ReadFile("functionalities/ApiEndpoints.php")
		file, err := os.Create(exPath + "/Functionality/ApiEndpoints.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("API ENDPOINTS CLASS CREATED")
		}
		file.Close()
	case "cf", "fields", "field", "custom-fields", "custom-field":
		original_text, _ := content.ReadFile("functionalities/CustomFields.php")
		file, err := os.Create(exPath + "/Functionality/CustomFields.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("CUSTOM FIELDS CLASS CREATED")
			script.Exec("composer require htmlburger/carbon-fields").Stdout()
		}
		file.Close()
	case "cpt", "custom-post-types", "custom-post-type", "post-type", "post-types":
		original_text, _ := content.ReadFile("functionalities/CustomPostTypes.php")
		file, err := os.Create(exPath + "/Functionality/CustomPostTypes.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("CPT CLASS CREATED")
		}
		file.Close()
	case "post-action", "post-actions", "post":
		original_text, _ := content.ReadFile("functionalities/PostActions.php")
		file, err := os.Create(exPath + "/Functionality/PostActions.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("POST ACTIONS CREATED")
		}
		file.Close()
	case "routes", "route":
		original_text, _ := content.ReadFile("functionalities/Routes.php")
		file, err := os.Create(exPath + "/Functionality/Routes.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("ROUTES CLASS CREATED")
			script.Exec("composer require joanrodas/plubo-routes").Stdout()
		}
		file.Close()
	case "roles", "role":
		original_text, _ := content.ReadFile("functionalities/Roles.php")
		file, err := os.Create(exPath + "/Functionality/Roles.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("ROLES CLASS CREATED")
			script.Exec("composer require joanrodas/plubo-roles").Stdout()
		}
		file.Close()
	case "shortcode", "shortcodes":
		original_text, _ := content.ReadFile("functionalities/Shortcodes.php")
		file, err := os.Create(exPath + "/Functionality/Shortcodes.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("SHORTCODES CLASS CREATED")
		}
		file.Close()
	case "tax", "taxonomy", "taxonomies":
		original_text, _ := content.ReadFile("functionalities/Taxonomies.php")
		file, err := os.Create(exPath + "/Functionality/Taxonomies.php")
		new_text = namespace_file(exPath, string(original_text))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println("TAXONOMIES CLASS CREATED")
		}
		file.Close()
	default:
		original_text, _ := content.ReadFile("functionalities/BaseFunctionality.php")
		to_replace, _ := script.Echo(functionality).Exec("tr -d '\n'").Exec("awk 'BEGIN{FS=\"\";RS=\"-\";ORS=\"\"} {$0=toupper(substr($0,1,1)) substr($0,2)} 1'").String()

		file, err := os.Create(exPath + "/Functionality/" + to_replace + ".php")
		new_text := namespace_file(exPath, strings.Replace(string(original_text), "BaseFunctionality", to_replace, -1))
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(new_text))
			fmt.Println(functionality + " CLASS CREATED")
		}
		file.Close()
	}
}

func component(name string) {
	original_text, _ := content.ReadFile("components/BaseComponent.php")
	exPath, _ := script.Exec("pwd").Exec("tr -d '\n'").String()
	to_replace, _ := script.Echo(name).Exec("tr -d '\n'").Exec("awk 'BEGIN{FS=\"\";RS=\"-\";ORS=\"\"} {$0=toupper(substr($0,1,1)) substr($0,2)} 1'").String()

	file, err := os.Create(exPath + "/Components/" + to_replace + ".php")
	new_text := namespace_file(exPath, strings.Replace(string(original_text), "BaseComponent", to_replace, -1))
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(string(new_text))
		fmt.Println("COMPONENT " + to_replace + " CREATED")
	}
	file.Close()
}

func utils(name string) {
	exPath, _ := script.Exec("pwd").Exec("tr -d '\n'").String()

	file, err := os.Create(exPath + "/Utils/" + name + ".php")
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString("")
		fmt.Println("UTILS " + name + " CREATED")
	}
	file.Close()
}

func lib(name string) {
	exPath, _ := script.Exec("pwd").Exec("tr -d '\n'").String()

	switch name {
	case "alpine":
		script.Exec("yarn add alpinejs").Stdout()
		original_text, _ := content.ReadFile("alpine/alpine.ts")
		file, err := os.Create(exPath + "/resources/scripts/components/alpine.ts")
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(original_text))
			fmt.Println("ADD ALPINE")
		}
		file.Close()
	case "tailwind":
		script.Exec("yarn add tailwindcss").Stdout()
		original_text, _ := content.ReadFile("tailwind/_tailwind.scss")
		file, err := os.Create(exPath + "/resources/styles/components/_tailwind.scss")
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(original_text))
			fmt.Println("ADD TAILWIND IMPORT")
		}
		file.Close()

		original_text_config, _ := content.ReadFile("tailwind/tailwind.config.cjs")
		file_config, err_config := os.Create(exPath + "/tailwind.config.cjs")
		if err_config != nil {
			fmt.Println(err_config)
		} else {
			file_config.WriteString(string(original_text_config))
			fmt.Println("ADD TAILWIND CONFIG")
		}
		file.Close()
	case "env":
		script.Exec("composer require vlucas/phpdotenv").Stdout()
		script.Exec("touch .env").Stdout()
		fmt.Println("ADD ENV")
	case "react":
		fmt.Println("ADD REACT")
	case "vue":
		fmt.Println("ADD VUE")
	default:
		color.Cyan("\nTo the rational mind, nothing is inexplicable; only unexplained.\n\n")
	}
}

func namespace_file(path, file_contents string) (new_text string) {
	to_replace, _ := script.Exec("basename " + path).Exec("tr -d '\n'").Exec("awk 'BEGIN{FS=\"\";RS=\"-\";ORS=\"\"} {$0=toupper(substr($0,1,1)) substr($0,2)} 1'").String()
	new_text, _ = script.Echo(string(file_contents)).Replace("PluginPlaceholder", to_replace).String()
	return
}

func namespace_project(path string) {
	to_replace, _ := script.Exec("basename " + path).Exec("tr -d '\n'").String()
	to_replace_mayus, _ := script.Exec("basename " + path).Exec("tr -d '\n'").Exec("awk 'BEGIN{FS=\"\";RS=\"-\";ORS=\"\"} {$0=toupper($0)} 1'").String()
	to_replace_pascal, _ := script.Exec("basename " + path).Exec("tr -d '\n'").Exec("awk 'BEGIN{FS=\"\";RS=\"-\";ORS=\"\"} {$0=toupper(substr($0,1,1)) substr($0,2)} 1'").String()
	os.Rename("./plugin-placeholder.php", "./"+to_replace+".php")
	os.Rename("./languages/plugin-placeholder.pot", "./languages/"+to_replace+".pot")

	err := filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !!fi.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*.php", fi.Name())

		if err != nil {
			panic(err)
		}

		if matched {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			newContents := strings.Replace(string(read), "plugin-placeholder", to_replace, -1)
			newContents = strings.Replace(newContents, "PLUGIN_PLACEHOLDER", to_replace_mayus, -1)
			newContents = strings.Replace(newContents, "PluginPlaceholder", to_replace_pascal, -1)

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}

		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	script.Exec("composer install").Stdout()
	return
}

func cleanup_files() {
	os.Remove("./.github/workflows/on-template.yml")
	os.Remove("./renovate.json")
	os.Remove("./README.md")
	os.Remove("./SECURITY.md")
	os.Remove("./CONTRIBUTING.md")
	os.Remove("./CODE_OF_CONDUCT.md")
	return
}

func init_template() {
	current_repo := os.Getenv("GITHUB_REPOSITORY")
	official_repo := "joanrodas/plubo"
	exPath, _ := script.Exec("pwd").Exec("tr -d '\n'").String()

	if current_repo == official_repo {
		fmt.Println("Not using template")
		return
	}

	if current_repo == "" {
		fmt.Println("Unknown Github Repo")
	}

	namespace_project(exPath)

	to_replace_pascal, _ := script.Exec("basename " + exPath).Exec("tr -d '\n'").Exec("awk 'BEGIN{FS=\"\";RS=\"-\";ORS=\"\"} {$0=toupper(substr($0,1,1)) substr($0,2)} 1'").String()

	//COMPOSER.JSON
	composer_name := strings.ToLower(current_repo)
	composer, _ := script.File("./composer.json").String()
	var composer_json map[string]interface{}
	var autoload_json map[string]interface{}
	json.Unmarshal([]byte(composer), &composer_json)
	composer_json["name"] = composer_name
	autoload, _ := json.MarshalIndent(composer_json["autoload"], "", "")
	changed_autoload := strings.Replace(string(autoload), "PluginPlaceholder", to_replace_pascal, -1)
	json.Unmarshal([]byte(changed_autoload), &autoload_json)
	composer_json["autoload"] = autoload_json
	composer_json["description"] = "An amazing plugin made with PLUBO"
	composer_json["type"] = "wordpress-plugin"
	composer_data, _ := json.MarshalIndent(composer_json, "", "\t")
	script.Echo(string(composer_data)).WriteFile("composer.json")

	//PACKAGE.JSON
	package_name := ("@" + composer_name)
	pack, _ := script.File("./package.json").String()
	var package_json map[string]interface{}
	json.Unmarshal([]byte(pack), &package_json)
	package_json["name"] = package_name
	package_data, _ := json.MarshalIndent(package_json, "", "\t")
	script.Echo(string(package_data)).WriteFile("package.json")

	cleanup_files()
}
