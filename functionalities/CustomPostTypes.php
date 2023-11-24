<?php

namespace PluginPlaceholder\Functionality;

class CustomPostTypes
{

    protected $plugin_name;
    protected $plugin_version;

    public function __construct($plugin_name, $plugin_version)
    {
        $this->plugin_name = $plugin_name;
        $this->plugin_version = $plugin_version;

        add_action('init', [$this, 'register_post_types']);
    }

    public function register_post_types()
    {
        // register_post_type('plubo', [
        //     'public'    => true,
        //     'show_in_rest' => true,
        //     'label'     => __('Plubo', 'plugin-placeholder'),
        //     'menu_icon' => 'dashicons-book',
        //     'capability_type'    => 'post',
        //     'has_archive'        => true,
        //     'hierarchical'       => false,
        //     'menu_position'      => 20,
        //     'supports'           => ['title', 'editor', 'author', 'thumbnail'],
        // ]);
    }
}
