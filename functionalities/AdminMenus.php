<?php

namespace PluginPlaceholder\Functionality;

use PluginPlaceholder\Includes\BladeLoader;

class AdminMenus
{

    protected $plugin_name;
    protected $plugin_version;

    private $blade;

    public function __construct($plugin_name, $plugin_version)
    {
        $this->plugin_name = $plugin_name;
        $this->plugin_version = $plugin_version;
        $this->blade = BladeLoader::getInstance();

        add_action('admin_menu', [$this, 'add_admin_menus']);
    }

    public function add_admin_menus()
    {
        // add_menu_page(
        //     __('PAGE TITLE', 'plugin-placeholder'),
        //     __('MENU TITLE', 'plugin-placeholder'),
        //     'manage_options',
        //     'plugin-placeholder',
        //     function () {
        //         echo $this->blade->template('settings');
        //     },
        //     'dashicons-admin-settings',
        //     6
        // );
    }
}
