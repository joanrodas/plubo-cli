<?php

namespace PluginPlaceholder\Functionality;

class BaseFunctionality
{

    protected $plugin_name;
    protected $plugin_version;

    public function __construct($plugin_name, $plugin_version)
    {
        $this->plugin_name = $plugin_name;
        $this->plugin_version = $plugin_version;

        add_action('init', [$this, 'example_function']);
    }

    public function example_function()
    {
    }
}
