components {
  id: "bg"
  component: "/game/levels/bg.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "offset"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "mult"
    value: "0.1"
    type: PROPERTY_TYPE_NUMBER
  }
  property_decls {
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/main.atlas\"\ndefault_animation: \"stars\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ADD\n"
  position {
    x: 120.0
    y: 8.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "tile_set: \"/main/main.atlas\"\ndefault_animation: \"stars\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ADD\n"
  position {
    x: 360.0
    y: 8.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
